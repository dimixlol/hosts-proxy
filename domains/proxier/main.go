package proxier

import (
	"context"
	"fmt"
	"github.com/dimixlol/knowyourwebsite/adapters/cache"
	"github.com/dimixlol/knowyourwebsite/adapters/database"
	"github.com/dimixlol/knowyourwebsite/config"
	"github.com/dimixlol/knowyourwebsite/domains/proxier/entrypoints"
	"github.com/dimixlol/knowyourwebsite/utils"
	"net/http"
)

func NewHTTPRequester(ctx context.Context) *http.Server {
	engine := utils.NewEngine()
	persister := database.NewDatabasePersister(ctx)
	cacheManager := cache.NewCacheManager(ctx, persister)
	entrypoints.NewHTTPEntrypoint(ctx, cacheManager, engine)
	return &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Configuration.Host, "8081"),
		Handler: engine,
	}
}
