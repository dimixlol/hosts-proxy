package usecases

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/dimixlol/hosts-proxy/logging"
	"github.com/dimixlol/hosts-proxy/pkg/compressor"
	"github.com/dimixlol/hosts-proxy/ports"
	"github.com/dimixlol/hosts-proxy/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
)

var (
	modifiers = []Modifier{embeddedUrlReplacer(), hostReplacer(), redirectReplacer()}
)

type Modifier func(c *gin.Context, url ports.URL, body string, response *http.Response) (string, error)

func handleError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": err,
	})
}

func NewRequestProxier(cache ports.CacheManager, dataCompressor compressor.Compressor) gin.HandlerFunc {
	return func(c *gin.Context) {
		host := c.Request.Host
		slug := utils.GetSlugFromHost(host)
		urlFromCache, err := cache.GetUrlBySlug(slug)

		if err != nil {
			handleError(c, err)
			return
		}

		fmt.Println(host, slug, urlFromCache)
		remote, err := url.Parse(httpSchema + urlFromCache.GetIP())

		if err != nil {
			handleError(c, err)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.ModifyResponse = modifyResponse(c, modifiers, dataCompressor, urlFromCache)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = urlFromCache.GetHost()
			req.URL.Scheme = strings.ReplaceAll(httpSchema, "://", "")
			req.URL.Host = urlFromCache.GetIP()
			req.URL.Path = c.Request.URL.Path
		}
		proxy.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // Disables certificate verification.
			},
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func modifyResponse(c *gin.Context, modifiers []Modifier, dataCompressor compressor.Compressor, url ports.URL) func(*http.Response) error {
	return func(res *http.Response) error {
		logger := logging.GetLogger(c)
		encoding := res.Header.Get("Content-Encoding")

		body := dataCompressor.Decompress(encoding, res.Body)

		for i, modifier := range modifiers {
			var err error
			body, err = modifier(c, url, body, res)
			if err != nil {
				logger.Errorf(c, "modifier %d failed while modifying response: %s", i, err)
				return err
			}
		}

		bodyBytes := dataCompressor.Compress(encoding, body)
		bodyContentLength := len(bodyBytes)
		res.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		res.ContentLength = int64(bodyContentLength)
		res.Header.Set("Content-Length", strconv.Itoa(bodyContentLength))

		return nil
	}
}
