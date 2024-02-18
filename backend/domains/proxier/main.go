package proxier

import (
	"context"
	"fmt"
	"github.com/dimixlol/hosts-proxy/adapters/cache"
	"github.com/dimixlol/hosts-proxy/adapters/storage"
	"github.com/dimixlol/hosts-proxy/config"
	"github.com/dimixlol/hosts-proxy/domains/proxier/entrypoints"
	"github.com/dimixlol/hosts-proxy/pkg/compressor"
	"github.com/dimixlol/hosts-proxy/utils"
	"net/http"
)

func NewHTTPRequester(ctx context.Context) *http.Server {
	engine := utils.NewEngine()
	db := storage.NewDatabaseStorage(ctx)
	cacheManager := cache.NewCacheManager(ctx, db)
	dataCompressor := compressor.New(ctx)
	entrypoints.NewHTTPEntrypoint(ctx, cacheManager, dataCompressor, engine)
	return &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Configuration.Proxier.Host, config.Configuration.Proxier.Port),
		Handler: engine,
	}
}
