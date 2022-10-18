package cache

import (
	"errors"

	"github.com/zhangrt/voyager1_core/constant"
)

//实现Cache的简单工厂
type CacheFactory struct{}

func (factory *CacheFactory) Create(cacheType string) (Cacher, error) {
	if cacheType == constant.REDIS_STANDALONE {
		return &RedisCache{}, nil
	}

	if cacheType == constant.REDIS_CLUSTER {
		return &RedisClusterCache{}, nil
	}
	return nil, errors.New("error cache type")
}
