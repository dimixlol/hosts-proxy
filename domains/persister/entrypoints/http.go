package entrypoints

import (
	"context"
	"fmt"
	"github.com/dimixlol/knowyourwebsite/config"
	"github.com/dimixlol/knowyourwebsite/domains/persister/usecases"
	"github.com/dimixlol/knowyourwebsite/logging"
	"github.com/dimixlol/knowyourwebsite/models"
	"github.com/dimixlol/knowyourwebsite/ports"
	"github.com/dimixlol/knowyourwebsite/static"
	"github.com/dimixlol/knowyourwebsite/utils"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
	"net/http"
	"strconv"
)

var (
	newPersistentSiteSpec = []fizz.OperationOption{
		fizz.ID("CreatePersistentSite"),
		fizz.Summary("Create Site"),
		fizz.Description("Create persistent site mapping with host to ip"),
		fizz.StatusDescription("Successful Response"),
		fizz.Response(strconv.Itoa(http.StatusBadRequest), "Invalid IP/Host", nil, nil, &utils.Response{Status: http.StatusBadRequest, Err: "invalid data passed"}),
	}
	getURLBySlugSpec = []fizz.OperationOption{
		fizz.ID("GetURLBySlug"),
		fizz.Summary("Get URL By Slug"),
		fizz.Description("Get URL By Slug"),
		fizz.StatusDescription("Successful Response"),
		fizz.Response(strconv.Itoa(http.StatusBadRequest), "Invalid slug", nil, nil, &utils.Response{Status: http.StatusBadRequest, Err: "invalid slug"}),
		fizz.Response(strconv.Itoa(http.StatusNotFound), "URL not found", nil, nil, &utils.Response{Status: http.StatusNotFound, Err: "url not found"}),
	}
)

type (
	newPersistentSiteRequest struct {
		Host string `json:"host" description:"fqdn" binding:"required"`
		IP   string `json:"ip" description:"ipv4 address" binding:"required"`
	}

	newPersistentSiteResponse struct {
		IP   string `json:"ip"`
		Host string `json:"host"`
		Slug string `json:"slug"`
	}
	getURLBySlugRequest struct {
		Slug string `json:"slug" binding:"required" example:"abasicslug"`
	}
)

func getOpenApiInfo() *openapi.Info {
	return &openapi.Info{
		Title:       config.Configuration.API.Title,
		Description: config.Configuration.API.Description,
		Version:     config.Configuration.Version,
		Contact: &openapi.Contact{
			Name:  config.Configuration.API.Contact.Name,
			URL:   config.Configuration.API.Contact.URL,
			Email: config.Configuration.API.Contact.Email,
		},
		License: &openapi.License{
			Name: config.Configuration.API.License.Name,
			URL:  config.Configuration.API.License.URL,
		},
		TermsOfService: "string",
		XLogo: &openapi.XLogo{
			URL:             config.Configuration.API.Logo.URL,
			BackgroundColor: config.Configuration.API.Logo.Color,
			AltText:         config.Configuration.API.Logo.AltText,
			Href:            config.Configuration.API.Logo.Href,
		},
	}
}

func NewHTTPEntrypoint(ctx context.Context, persister ports.Persister, engine *fizz.Fizz) {
	api := engine.Group(fmt.Sprintf("/api/%s/persister", config.Configuration.API.Version), "persister", "Persister API")
	api.POST("create/", newPersistentSiteSpec, newPersistentSite(ctx, persister))
	api.POST("get/", getURLBySlugSpec, getURLBySlug(ctx, persister))
	engine.GET("/api/v1/swagger.json", nil, engine.OpenAPI(getOpenApiInfo(), "json"))
	engine.GET("/api/v1/swagger.yaml", nil, engine.OpenAPI(getOpenApiInfo(), "yaml"))
	engine.Engine().StaticFS("/api/swagger", http.FS(static.SwaggerAssets))
	engine.Engine().StaticFS("/api/redoc", http.FS(static.RedocAssets))
	engine.GET("/", nil, func(c *gin.Context) { c.Redirect(http.StatusMovedPermanently, "/api/redoc") })
}

func newPersistentSite(ctx context.Context, persister ports.Persister) gin.HandlerFunc {
	logger := logging.GetLogger(ctx)
	return tonic.Handler(
		func(c *gin.Context, req *newPersistentSiteRequest) (*newPersistentSiteResponse, error) {
			ip := models.NewIP(req.IP)
			if ip == nil {
				logger.Errorf(ctx, "invalid ip `%s` passed", req.IP)
				return nil, errors.NewBadRequest(nil, "invalid ip")
			}
			host, err := models.NewHost(req.Host)
			if err != nil {
				logger.Errorf(ctx, "invalid host `%s` passed: %s", req.Host, err)
				return nil, errors.NewBadRequest(err, "invalid host")
			}
			url := usecases.NewPersistentSite(ctx, persister, host, ip)
			if url == nil {
				panic("err")
			}
			return &newPersistentSiteResponse{IP: url.GetIP(), Host: url.GetHost(), Slug: url.GetSlug()}, nil
		},
		http.StatusOK,
	)
}

func getURLBySlug(ctx context.Context, persister ports.Persister) gin.HandlerFunc {
	logger := logging.GetLogger(ctx)
	return tonic.Handler(
		func(c *gin.Context, req *getURLBySlugRequest) (*newPersistentSiteResponse, error) {
			if uint(len(req.Slug)) != config.Configuration.SlugLength {
				logger.Errorf(ctx, "invalid slug length `%d` passed", len(req.Slug))
				return nil, errors.NewBadRequest(nil, "invalid slug")
			}
			url, err := persister.GetURLBySlug(req.Slug)
			if err != nil {
				logger.Errorf(ctx, "error while getting url by slug `%s`: %s", req.Slug, err)
				return nil, err
			}
			if url == nil {
				return nil, errors.NewNotFound(nil, "url not found")
			}
			return &newPersistentSiteResponse{IP: url.GetIP(), Host: url.GetHost(), Slug: url.GetSlug()}, nil
		},
		http.StatusOK,
	)
}
