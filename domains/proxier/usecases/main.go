package usecases

import (
	"bytes"
	"crypto/tls"
	"github.com/dimixlol/knowyourwebsite/pkg/compress"
	"github.com/dimixlol/knowyourwebsite/ports"
	"github.com/dimixlol/knowyourwebsite/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
)

var compressor compress.Compressor

func init() {
	compressor = compress.NewCompressor()
}

func NewRequestProxier(cache ports.CacheManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		host := c.Request.Host
		slug := utils.GetSlugFromHost(host)
		getUrlFromCache := cache.GetUrlBySlug(slug)

		remote, err := url.Parse("https://" + getUrlFromCache.GetIP())
		if err != nil {
			panic(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.ModifyResponse = modifier(c, getUrlFromCache.GetHost())
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = getUrlFromCache.GetHost()
			req.URL.Scheme = "https"
			req.URL.Host = getUrlFromCache.GetIP()
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

func modifier(c *gin.Context, oldHost string) func(*http.Response) error {
	return func(res *http.Response) error {
		encoding := res.Header.Get("Content-Encoding")
		body := compressor.Decompress(encoding, res.Body)
		body = strings.ReplaceAll(body, oldHost, c.Request.Host)

		bodyBytes := compressor.Compress(encoding, body)
		bodyContentLength := len(bodyBytes)

		writer := io.NopCloser(bytes.NewBuffer(bodyBytes))
		res.Body = writer
		res.ContentLength = int64(bodyContentLength)
		res.Header.Set("Content-Length", strconv.Itoa(bodyContentLength))
		return nil
	}
}
