package usecases

import (
	"github.com/dimixlol/hosts-proxy/config"
	"github.com/dimixlol/hosts-proxy/ports"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	httpSchema  = "http://"
	httpsSchema = "https://"
)

func embeddedUrlReplacer() Modifier {
	return func(c *gin.Context, url ports.URL, body string, response *http.Response) (string, error) {
		var (
			host   = c.Request.Host
			schema = httpsSchema
			port   string
		)

		if parts := strings.Split(c.Request.Host, ":"); len(parts) > 1 {
			port = ":" + config.Configuration.Proxier.Port
		}

		body = strings.ReplaceAll(body, "https://"+url.GetHost(), schema+host+port)

		return body, nil
	}
}

func hostReplacer() Modifier {
	return func(c *gin.Context, url ports.URL, body string, response *http.Response) (string, error) {
		return strings.ReplaceAll(body, url.GetHost(), c.Request.Host), nil
	}
}

func redirectReplacer() Modifier {
	return func(c *gin.Context, url ports.URL, body string, response *http.Response) (string, error) {
		if location := response.Header.Get("Location"); location != "" {
			response.Header.Set("Location", strings.ReplaceAll(location, url.GetHost(), c.Request.Host))
		}
		return body, nil
	}
}
