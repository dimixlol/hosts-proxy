package usecases

import (
	"github.com/dimixlol/knowyourwebsite/config"
	"github.com/dimixlol/knowyourwebsite/domains/proxier/ports"
	appPorts "github.com/dimixlol/knowyourwebsite/ports"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	httpSchema  = "http://"
	httpsSchema = "https://"
)

func embedUrlReplacer() ports.Modifier {
	return func(c *gin.Context, url appPorts.URL, body string, response *http.Response) (string, error) {
		var (
			host   = c.Request.Host
			schema = httpSchema
			port   string
		)

		if config.Configuration.Proxier.TLS {
			schema = httpsSchema
			if parts := strings.Split(c.Request.Host, ":"); len(parts) > 1 {
				port = ":" + config.Configuration.Proxier.Port
			}
		}

		body = strings.ReplaceAll(body, "https://"+url.GetHost(), schema+host+port)

		return body, nil
	}
}

func hostReplacer() ports.Modifier {
	return func(c *gin.Context, url appPorts.URL, body string, response *http.Response) (string, error) {
		return strings.ReplaceAll(body, url.GetHost(), c.Request.Host), nil
	}
}

func redirectReplacer() ports.Modifier {
	return func(c *gin.Context, url appPorts.URL, body string, response *http.Response) (string, error) {
		if location := response.Header.Get("Location"); location != "" {
			response.Header.Set("Location", strings.ReplaceAll(location, url.GetHost(), c.Request.Host))
		}
		return body, nil
	}
}
