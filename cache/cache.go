package cache

import (
	"time"
)

type Cache struct {
	CacheType        string
	ConnectionString string
}

type Cacher interface {
	Connect(constr string)
	Close() bool
	Get(key string) string
	Set(key string, value string, exp time.Duration) string
	Del(key string) bool
	HGet(key string, field string) string
	HSet(key string, field string, value string) string
	Hmget(key string, fields []string) []string
	Hmset(key string, maps map[string]string) string
	Hexists(key string, field string) bool
	Lpush(key string, value string) string
	Rpush(key string, value string) string
	Lpop(key string) string
	Rpop(key string) string
	Llen(key string) int64
	Linsert(key string, operateType string, existsValue string, value string) string
	Lrang(key string, start int64, end int64) []string
}
