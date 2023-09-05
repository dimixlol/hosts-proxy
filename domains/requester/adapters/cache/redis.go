package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/dimixlol/knowyourwebsite/config"
	"github.com/dimixlol/knowyourwebsite/domains/persister/adapters/database"
	"github.com/dimixlol/knowyourwebsite/domains/persister/models"
	persisterPorts "github.com/dimixlol/knowyourwebsite/domains/persister/ports"
	"github.com/dimixlol/knowyourwebsite/domains/requester/ports"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type redisCacheManager struct {
	ctx       context.Context
	cache     *cache.Cache
	persister persisterPorts.Persister
}

func (r *redisCacheManager) GetUrlBySlug(slug string) persisterPorts.URL {
	var cached *models.URL
	var err error

	err = r.cache.Get(r.ctx, slug, &cached)
	if err != nil && !errors.Is(err, cache.ErrCacheMiss) {
		panic(err)
	}

	if cached == nil {
		cached, err := r.persister.GetURLBySlug(slug)
		err = r.cache.Set(&cache.Item{
			Ctx:   r.ctx,
			Key:   slug,
			Value: cached,
			TTL:   config.Configuration.Cache.TTL,
		})
		if err != nil {
			panic(err)
		}
	}
	return cached
}

func newRedisCacheManager(ctx context.Context) ports.CacheManager {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Configuration.Cache.Host, config.Configuration.Cache.Port),
		Password: config.Configuration.Cache.Password,
		DB:       config.Configuration.Cache.DB,
	})

	rc := cache.New(&cache.Options{
		Redis:      rdb,
		LocalCache: cache.NewTinyLFU(config.Configuration.Cache.Size, config.Configuration.Cache.TTL),
	})

	return &redisCacheManager{
		ctx,
		rc,
		database.NewDatabasePersister(ctx),
	}
}

func NewCacheManager(ctx context.Context) ports.CacheManager {
	return newRedisCacheManager(ctx)
}
