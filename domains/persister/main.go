package persister

import (
	"context"
	"fmt"
	"github.com/dimixlol/knowyourwebsite/adapters/database"
	"github.com/dimixlol/knowyourwebsite/config"
	"github.com/dimixlol/knowyourwebsite/domains/persister/entrypoints"
	"github.com/dimixlol/knowyourwebsite/utils"
	"net/http"
)

func NewHTTPPersister(ctx context.Context) *http.Server {
	engine := utils.NewEngine()
	db := database.NewDatabasePersister(ctx)
	entrypoints.NewHTTPEntrypoint(ctx, db, engine)
	return &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Configuration.Host, config.Configuration.Port),
		Handler: engine,
	}
}
