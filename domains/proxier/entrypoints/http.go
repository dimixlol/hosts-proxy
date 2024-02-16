package entrypoints

import (
	"context"
	"github.com/dimixlol/hosts-proxy/domains/proxier/usecases"
	"github.com/dimixlol/hosts-proxy/pkg/compressor"
	"github.com/dimixlol/hosts-proxy/ports"
	"github.com/wI2L/fizz"
)

func NewHTTPEntrypoint(ctx context.Context, cache ports.CacheManager, dataCompressor compressor.Compressor, engine *fizz.Fizz) {
	proxier := usecases.NewRequestProxier(cache, dataCompressor)
	engine.Engine().Any("/*any", proxier)
}
