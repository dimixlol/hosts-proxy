package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/dimixlol/knowyourwebsite/config"
	"github.com/dimixlol/knowyourwebsite/logging"
	"github.com/dimixlol/knowyourwebsite/models"
	"github.com/dimixlol/knowyourwebsite/ports"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type redisCacheManager struct {
	ctx       context.Context
	cache     *cache.Cache
	persister ports.Persister
	logger    *logging.Logger
}

func (r *redisCacheManager) GetUrlBySlug(slug string) (ports.URL, error) {
	var cached *models.URL
	var err error

	err = r.cache.Get(r.ctx, slug, &cached)
	if err != nil && !errors.Is(err, cache.ErrCacheMiss) {
		r.logger.Errorf(r.ctx, "error while getting url from cache: %s", err.Error())
		return nil, err
	}

	if cached == nil {
		cached, err := r.persister.GetURLBySlug(slug)
		if err != nil {
			r.logger.Errorf(r.ctx, "error while getting url from persister to cache: %s", err.Error())
			return nil, err
		}
		if cached == nil {
			r.logger.Errorf(r.ctx, "url not found in persister")
			return nil, nil
		}
		err = r.cache.Set(&cache.Item{
			Ctx:   r.ctx,
			Key:   slug,
			Value: cached,
			TTL:   config.Configuration.Cache.TTL,
		})
		if err != nil {
			r.logger.Errorf(r.ctx, "error while setting url to cache: %s", err.Error())
		}
	}
	return cached, nil
}

func newRedisCacheManager(ctx context.Context, persister ports.Persister) ports.CacheManager {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Configuration.Cache.Host, config.Configuration.Cache.Port),
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
		persister,
		logging.GetLogger(ctx),
	}
}

func NewCacheManager(ctx context.Context, persister ports.Persister) ports.CacheManager {
	return newRedisCacheManager(ctx, persister)
}