package cache

import (
	"errors"
)

//实现Cache的简单工厂
type CacheFactory struct{}

func (factory *CacheFactory) Create(cacheType string) (Cacher, error) {
	if cacheType == "G_REDIS_STANDALONE" {
		return &RedisCache{}, nil
	}

	if cacheType == "G_REDIS_CLUSTER" {
		return &RedisClusterCache{}, nil
	}
	return nil, errors.New("error cache type")
}
