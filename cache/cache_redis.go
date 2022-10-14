package cache

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/zhangrt/voyager1_core/global"
)

//实现具体的Cache：RedisCache
type RedisCache struct {
}

var redis_Cache *redis.Client

// func getRedis() *redis.Client {
// 	if redis_Cache == nil {
// 		lock.Lock()
// 		defer lock.Unlock()
// 		if redis_Cache == nil {
// 			redis_Cache = createRedis()
// 		}
// 	}
// 	return redis_Cache
// }

func createRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     global.G_CONFIG.Cache.Addr,     // redis地址
		Password: global.G_CONFIG.Cache.Password, // redis密码，没有设置，则留空
		DB:       0,                              // 使用默sss数据库
	})
	_, err := client.Ping(ctx).Result()
	return client, err
}

func (redisCache *RedisCache) InitCache() error {
	var err error
	if redis_Cache == nil {
		lock.Lock()
		defer lock.Unlock()
		if redis_Cache == nil {
			redis_Cache, err = createRedis()
		}
	}
	return err
}

func Close() bool {
	redis_Cache.Close()
	return true
}

func (redisCache *RedisCache) Get(key string) (string, error) {
	result, err := redis_Cache.Get(ctx, key).Result()
	return result, err
}

func (redisCache *RedisCache) Set(key string, value interface{}) error {
	err := redis_Cache.Set(ctx, key, value, time.Second*0).Err()
	return err
}

func (redisCache *RedisCache) LPush(key string, value ...interface{}) error {
	err := redis_Cache.LPush(ctx, key, value).Err()
	return err
}

func (redisCache *RedisCache) RPush(key string, value ...interface{}) error {
	err := redis_Cache.RPush(ctx, key, value).Err()
	return err
}

func (redisCache *RedisCache) LPop(key string) (interface{}, error) {
	result, err := redis_Cache.LPop(ctx, key).Result()
	return result, err
}

func (redisCache *RedisCache) RPop(key string) (interface{}, error) {
	result, err := redis_Cache.RPop(ctx, key).Result()
	return result, err
}

func (redisCache *RedisCache) BLPop(timeout time.Duration, key string) (interface{}, error) {
	result, err := redis_Cache.BLPop(ctx, timeout, key).Result()
	return result, err
}

func (redisCache *RedisCache) BRPop(timeout time.Duration, key string) (interface{}, error) {
	result, err := redis_Cache.BRPop(ctx, timeout, key).Result()
	return result, err
}

func (redisCache *RedisCache) LLen(key string) (int64, error) {
	result, err := redis_Cache.LLen(ctx, key).Result()
	return result, err
}

func (redisCache *RedisCache) LRange(key string, start, end int64) ([]string, error) {
	result, err := redis_Cache.LRange(ctx, key, start, end).Result()
	return result, err
}

func (redisCache *RedisCache) HSet(hashKey, key string, value interface{}) error {
	err := redis_Cache.HSet(ctx, hashKey, key, value).Err()
	return err
}

func (redisCache *RedisCache) HGet(hashKey, key string) (interface{}, error) {
	result, err := redis_Cache.HGet(ctx, hashKey, key).Result()
	return result, err
}

func (redisCache *RedisCache) HGetAll(hashKey string) (map[string]string, error) {
	result, err := redis_Cache.HGetAll(ctx, hashKey).Result()
	return result, err
}

func (redisCache *RedisCache) HDel(hashKey string, key ...string) error {
	err := redis_Cache.HDel(ctx, hashKey, key...).Err()
	return err
}

func (redisCache *RedisCache) HExists(hashKey, key string) (bool, error) {
	result, err := redis_Cache.HExists(ctx, hashKey, key).Result()
	return result, err
}

func (redisCache *RedisCache) SAdd(key string, values ...interface{}) error {
	err := redis_Cache.SAdd(ctx, key, values).Err()
	return err
}

func (redisCache *RedisCache) SCard(key string) (int64, error) {
	result, err := redis_Cache.SCard(ctx, key).Result()
	return result, err
}

func (redisCache *RedisCache) SMembers(key string) ([]string, error) {
	result, err := redis_Cache.SMembers(ctx, key).Result()
	return result, err
}

func (redisCache *RedisCache) SRem(key string, value interface{}) error {
	err := redis_Cache.SRem(ctx, key, value).Err()
	return err
}

func (redisCache *RedisCache) SPop(key string) (interface{}, error) {
	result, err := redis_Cache.SPop(ctx, key).Result()
	return result, err
}

func (redisCache *RedisCache) Expire(key string, duration time.Duration) error {
	err := redis_Cache.Expire(ctx, key, duration).Err()
	return err
}

func (redisCache *RedisCache) ExpireAt(key string, duration time.Time) error {
	err := redis_Cache.ExpireAt(ctx, key, duration).Err()
	return err
}

func (redisCache *RedisCache) TTL(key string) (time.Duration, error) {
	result, err := redis_Cache.TTL(ctx, key).Result()
	return result, err
}
