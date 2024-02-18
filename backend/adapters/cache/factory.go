package cache

import (
	"context"
	"github.com/dimixlol/hosts-proxy/ports"
)

func NewCacheManager(ctx context.Context, persister ports.Storage) ports.CacheManager {
	return NewRedisCacheManager(ctx, persister)
}
