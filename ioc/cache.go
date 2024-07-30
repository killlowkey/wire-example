package ioc

import "wire-example/pkg/cache"

func NewCache() cache.ICache {
	return cache.NewRedisCache()
}
