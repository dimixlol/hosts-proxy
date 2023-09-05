package http

import (
	"github.com/dimixlol/knowyourwebsite/domains/persister/models"
	"github.com/dimixlol/knowyourwebsite/domains/persister/ports"
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
)

var NewPersistentSiteSpec = []fizz.OperationOption{
	fizz.ID("CreatePersistentSite"),
	fizz.Summary("Create Site"),
	fizz.Description("Create persistent site mapping with host to ip"),
	fizz.StatusDescription("Successful Response"),
	fizz.Response("404", "transaction not found", nil, nil, 404),
}

type NewPersistentSiteRequest struct {
	Host string `json:"host" binding:"required"`
	IP   string `json:"ip" binding:"required"`
}

type NewPersistentSiteResponse struct {
	IP   string `json:"ip"`
	Host string `json:"host"`
	Slug string `json:"slug"`
}

func NewPersistentSite(db ports.Persister) gin.HandlerFunc {
	return tonic.Handler(
		func(c *gin.Context, req *NewPersistentSiteRequest) (*NewPersistentSiteResponse, error) {
			ip := models.NewIP(req.IP)
			if ip == nil {
				panic("err")
			}
			host := models.NewHost(req.Host)
			if host == nil {
				panic("err")
			}
			url := db.NewPersistentSite(host, ip)
			if url == nil {
				panic("err")
			}
			return &NewPersistentSiteResponse{IP: url.GetIP(), Host: url.GetHost(), Slug: url.GetSlug()}, nil
		},
		200,
	)
}
