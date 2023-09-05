package requester

import (
	"context"
	"github.com/dimixlol/knowyourwebsite/domains/requester/adapters/cache"
	"github.com/dimixlol/knowyourwebsite/domains/requester/usecases/http"
	"github.com/gin-gonic/gin"
)

func NewRequester(ctx context.Context) gin.HandlerFunc {
	manager := cache.NewCacheManager(ctx)
	return http.NewHandler(manager)
}
