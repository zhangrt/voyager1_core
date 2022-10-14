package cache

import "errors"

type cacheType int

const (
	redisCache cacheType = iota
	memCache
)

//实现Cache的简单工厂
type CacheFactory struct{}

func (factory *CacheFactory) Create(cacheType cacheType) (Cacher, error) {
	if cacheType == redisCache {
		return &RedisCache{}, nil
	}
	return nil, errors.New("error cache type")
}
