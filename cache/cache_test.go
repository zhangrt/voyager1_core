package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestCacheFactory_Create(t *testing.T) {
	cacheFactory := &CacheFactory{}

	redis, error := cacheFactory.Create(redisCache)
	if error != nil {
		t.Error(error)
	}

	redis.Connect("localhost:49153")
	// err := redis.Connect("localhost:49153")
	// if err != nil {
	// 	fmt.Printf("redis连接失败! err : %v\n", err)
	// 	return
	// }
	fmt.Println("redis连接成功!")

	// 存普通string类型，10分钟过期
	redis.Set("test:name", "张三", time.Minute*10)
	// 存hash数据
	//redis.HSet("test:class", "521", "42")
	// 存list数据
	//redis.Rpush("test:list", "1") // 向右边添加元素
	//redis.Lpush("test:list", "2") // 向左边添加元素
	// 存set数据
	//redis.SAdd("test:set", "apple")
	//redis.SAdd("test:set", "pear")

	fmt.Println(redis.Get("test:name"))

	redis.Close()
	// mem, error := cacheFactory.Create(mem)
	// if error != nil {
	// 	t.Error(error)
	// }
	// mem.Set("k1", "v1")

	// fmt.Println(mem.Get("k1"))
}
