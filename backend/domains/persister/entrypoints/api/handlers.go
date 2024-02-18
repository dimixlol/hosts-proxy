package api

import (
	"context"
	"fmt"
	"github.com/dimixlol/hosts-proxy/config"
	"github.com/dimixlol/hosts-proxy/domains/persister/usecases"
	"github.com/dimixlol/hosts-proxy/logging"
	"github.com/dimixlol/hosts-proxy/models"
	"github.com/dimixlol/hosts-proxy/ports"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
	"net/http"
)

func NewHTTPAPI(ctx context.Context, storage ports.Storage, engine *fizz.Fizz) {
	logger := logging.GetLogger(ctx)
	api := engine.Group(fmt.Sprintf("/api/%s/persister", config.Configuration.API.Version), "persister", "Persister API")
	api.POST("create", newPersistentSiteSpec, newPersistentSite(ctx, logger, storage))
	api.POST("get", getURLBySlugSpec, getURLBySlug(ctx, logger, storage))
}

func newPersistentSite(ctx context.Context, logger *logging.Logger, storage ports.Storage) gin.HandlerFunc {
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
			url := usecases.NewPersistentSite(ctx, storage, host, ip)
			if url == nil {
				panic("err")
			}
			return &newPersistentSiteResponse{IP: url.GetIP(), Host: url.GetHost(), Slug: url.GetSlug()}, nil
		},
		http.StatusOK,
	)
}

func getURLBySlug(ctx context.Context, logger *logging.Logger, storage ports.Storage) gin.HandlerFunc {
	return tonic.Handler(
		func(c *gin.Context, req *getURLBySlugRequest) (*newPersistentSiteResponse, error) {
			if uint(len(req.Slug)) != config.Configuration.SlugLength {
				logger.Errorf(ctx, "invalid slug length `%d` passed", len(req.Slug))
				return nil, errors.NewBadRequest(nil, "invalid slug")
			}
			url, err := storage.GetURLBySlug(req.Slug)
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
