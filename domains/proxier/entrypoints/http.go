package entrypoints

import (
	"context"
	"github.com/dimixlol/knowyourwebsite/domains/proxier/usecases"
	"github.com/dimixlol/knowyourwebsite/ports"
	"github.com/wI2L/fizz"
)

func NewHTTPEntrypoint(ctx context.Context, cache ports.CacheManager, engine *fizz.Fizz) {
	proxier := usecases.NewRequestProxier(cache)
	engine.Engine().Any("/*any", proxier)
}
