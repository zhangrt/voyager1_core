package cache

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/zhangrt/voyager1_core/global"
)

//var Cache Cacher

var lock = &sync.Mutex{} //创建互锁
var ctx = context.Background()

type Cacher interface {
	//初始化缓存
	InitCache() error
	//Get 获取缓存数据
	Get(key string) (string, error)
	//Set 设置数据 默认不过期
	Set(key string, value interface{}) error
	//LPush 从列表的头部进行添加
	LPush(key string, value ...interface{}) error
	//RPush 使用RPush命令往队列右边加入
	RPush(key string, value ...interface{}) error
	//LPop 取出并移除左边第一个元素
	LPop(key string) (interface{}, error)
	//RPop 取出并移除右边第一个元素
	RPop(key string) (interface{}, error)
	//BLPop 取出并移除左边第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
	BLPop(timeout time.Duration, key string) (interface{}, error)
	//BRPop 取出并移除右边第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
	BRPop(timeout time.Duration, key string) (interface{}, error)
	//获取数据长度
	LLen(key string) (int64, error)
	//获取数据列表
	LRange(key string, start, end int64) ([]string, error)
	//hash 适合存储结构
	HSet(hashKey, key string, value interface{}) error
	//get Hash
	HGet(hashKey, key string) (interface{}, error)
	//HGetAll 获取所以hash ,返回map
	HGetAll(hashKey string) (map[string]string, error)
	//HDel 删除一个或多个哈希表字段
	HDel(hashKey string, key ...string) error
	//HExists 查看哈希表的指定字段是否存在
	HExists(hashKey, key string) (bool, error)
	//添加Set
	SAdd(key string, values ...interface{}) error
	//SCard 获取集合的成员数
	SCard(key string) (int64, error)
	//SMembers 获取集合的所有成员
	SMembers(key string) ([]string, error)
	//SRem 移除集合里的某个元素
	SRem(key string, value interface{}) error
	//SPop 移除并返回set的一个随机元素(SET是无序的)
	SPop(key string) (interface{}, error)
	//Expire 给指定key 设置过期时间
	Expire(key string, duration time.Duration) error
	//ExpireAt 给指定Key 设置过期时间，时间格式为UNIX时间
	ExpireAt(key string, duration time.Time) error
	//TTL 获取key的生存时间
	TTL(key string) (time.Duration, error)
}

// func init() {
// 	fmt.Println("创建缓存工厂，根据配置创建对应缓存组件")
// 	cacheFactory := &CacheFactory{}

// 	redis, error := cacheFactory.Create(global.G_CONFIG.Cache.Options)
// 	fmt.Println("global.G_CONFIG.Cache.Options" + global.G_CONFIG.Cache.Options)
// 	//redis, error := cacheFactory.Create("G_REDIS_STANDALONE")
// 	if error != nil {
// 		fmt.Printf("缓存工厂创建! error : %v\n", error)
// 	}
// 	Cache = redis
// }

func CreateCache() Cacher {
	fmt.Println("创建缓存工厂，根据配置创建对应缓存组件")
	cacheFactory := &CacheFactory{}

	redis, error := cacheFactory.Create(global.G_CONFIG.Cache.Options)
	if error != nil {
		fmt.Printf("缓存工厂创建! error : %v\n", error)
	}
	error = redis.InitCache()
	if error != nil {
		fmt.Printf("缓存组件创建! error : %v\n", error)
	}
	return redis
}
