package cache

import (
	"context"
	"fmt"
)

type ICache interface {
	Get(ctx context.Context, key string) (any, error)
}

type RedisCache struct {
}

func NewRedisCache() ICache {
	return &RedisCache{}
}

func (rc *RedisCache) Get(ctx context.Context, key string) (any, error) {
	return fmt.Sprintf("value=%s", key), nil
}
