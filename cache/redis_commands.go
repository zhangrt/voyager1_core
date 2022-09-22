package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Redis_Cache *redis.Client
var ctx = context.Background()

func ConnRedis(constr string) {
	client := redis.NewClient(&redis.Options{
		Addr:     constr, // redis地址
		Password: "",     // redis密码，没有设置，则留空
		DB:       0,      // 使用默sss数据库
	})

	Redis_Cache = client
}

func Get(key string) string {
	val, err := Redis_Cache.Get(ctx, key).Result()
	// 判断查询是否出错
	if err != nil {
		panic(err)
	}
	return val
}

func Set(key string, value string) string {
	err := Redis_Cache.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
	return key
}

func Del(key string) bool {
	// 删除key
	err := Redis_Cache.Del(ctx, key).Err()
	if err != nil {
		panic(err)
	}
	return true
}
