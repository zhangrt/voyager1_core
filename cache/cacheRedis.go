package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

//实现具体的Cache：RedisCache
type RedisCache struct {
	//redis_Cache *redis.Client
	//data        map[string]string
}

var redis_Cache *redis.Client
var ctx = context.Background()

// func (redisCache *RedisCache) Connect(constr string) (err error) {
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     constr,    // redis地址
// 		Password: "redispw", // redis密码，没有设置，则留空
// 		DB:       0,         // 使用默sss数据库
// 	})
// 	redis_Cache = client
// 	_, err = client.Ping(ctx).Result()
// 	return
// }

func (redisCache *RedisCache) Connect(constr string) {
	client := redis.NewClient(&redis.Options{
		Addr:     constr,    // redis地址
		Password: "redispw", // redis密码，没有设置，则留空
		DB:       0,         // 使用默sss数据库
	})
	redis_Cache = client
}

func (redisCache *RedisCache) Close() bool {
	redis_Cache.Close()
	return true
}

func (redisCache *RedisCache) Get(key string) string {
	val, err := redis_Cache.Get(ctx, key).Result()
	// 判断查询是否出错
	if err != nil {
		panic(err)
	}
	return val
}

func (redisCache *RedisCache) Set(key string, value string, exp time.Duration) string {
	err := redis_Cache.Set(ctx, key, value, exp).Err()
	if err != nil {
		panic(err)
	}
	return key
}

func (redisCache *RedisCache) Del(key string) bool {
	// 删除key
	value, err := redis_Cache.Del(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	if value == 1 {
		return true
	} else {
		return false
	}
}

func (redisCache *RedisCache) HGet(key string, field string) string {
	value, err := redis_Cache.HGet(ctx, key, field).Result()
	if err != nil {
		panic(err)
	}
	return value
}

func (redisCache *RedisCache) HSet(key string, field string, value string) string {
	err := redis_Cache.HSet(ctx, key, field, value).Err()
	if err != nil {
		panic(err)
	}
	return key
}

func (redisCache *RedisCache) Hmget(key string, fields []string) []string {
	keyLen := len(fields)
	values := make([]string, keyLen)
	for i := 0; i < len(fields); i++ {
		values[i] = redisCache.HGet(key, fields[i])
	}

	return values
}

func (redisCache *RedisCache) Hmset(key string, maps map[string]string) string {

	err := redis_Cache.HMSet(ctx, key, maps).Err()
	if err != nil {
		panic(err)
	}

	return key
}

func (redisCache *RedisCache) Hexists(key string, field string) bool {
	value, err := redis_Cache.HExists(ctx, key, field).Result()
	if err != nil {
		panic(err)
	}
	return value
}

func (redisCache *RedisCache) Lpush(key string, value string) string {
	// LPush支持一次插入任意个数据
	err := redis_Cache.LPush(ctx, key, value).Err()
	if err != nil {
		panic(err)
	}
	return key
}

func (redisCache *RedisCache) Rpush(key string, value string) string {
	// LPush支持一次插入任意个数据
	err := redis_Cache.RPush(ctx, key, value).Err()
	if err != nil {
		panic(err)
	}
	return key
}

func (redisCache *RedisCache) Lpop(key string) string {
	// Lpop从列表左边删除第一个数据，并返回删除的数据
	value, err := redis_Cache.LPop(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return value
}

func (redisCache *RedisCache) Rpop(key string) string {
	// Rpop从列表右边删除第一个数据，并返回删除的数据
	value, err := redis_Cache.RPop(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return value
}

func (redisCache *RedisCache) Llen(key string) int64 {
	val, err := redis_Cache.LLen(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return val
}

func (redisCache *RedisCache) Linsert(key string, operateType string, existsValue string, value string) string {
	switch operateType {
	case "before":
		err := redis_Cache.LInsert(ctx, key, "before", existsValue, value).Err()
		if err != nil {
			panic(err)
		}
	case "after":
		err := redis_Cache.LInsert(ctx, key, "after", existsValue, value).Err()
		if err != nil {
			panic(err)
		}
	}
	return key
}

func (redisCache *RedisCache) Lrang(key string, start int64, end int64) []string {
	values, err := redis_Cache.LRange(ctx, key, start, end).Result()
	if err != nil {
		panic(err)
	}
	return values
}