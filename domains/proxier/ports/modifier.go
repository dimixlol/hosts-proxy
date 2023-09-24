package ports

import (
	"github.com/dimixlol/knowyourwebsite/ports"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Modifier func(c *gin.Context, url ports.URL, body string, response *http.Response) (string, error)
