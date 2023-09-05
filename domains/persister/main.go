package persister

import (
	"context"
	"github.com/dimixlol/knowyourwebsite/domains/persister/adapters/database"
	"github.com/dimixlol/knowyourwebsite/domains/persister/handlers/http"
	"github.com/wI2L/fizz"
)

func NewPersister(ctx context.Context, fz *fizz.Fizz) {
	api := fz.Group("/persister", "persister", "Persister API")
	db := database.NewDatabasePersister(ctx)
	api.POST("create/", http.NewPersistentSiteSpec, http.NewPersistentSite(db))
}
