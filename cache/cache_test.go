package cache

import (
	"fmt"
	"testing"

	"github.com/zhangrt/voyager1_core/config"
	"github.com/zhangrt/voyager1_core/global"
)

//var Cache Cacher

func TestCacheFactory_Create(t *testing.T) {
	fmt.Println("初始化缓存组件所需配置信息")
	c := config.Cache{}
	c.Addr = "localhost:49153"
	c.Password = "redispw"
	c.Options = "G_REDIS_STANDALONE"
	global.G_CONFIG.Cache = c
	Cache := CreateCache() //在初始化配置参数后，创建缓存组件，将Cache可以保存在一个全局变量使用
	var err error
	//get set 操作
	err = Cache.Set("test:name", "li4")
	if err != nil {
		fmt.Printf("error : %v\n", err)
		return
	}
	s, err := Cache.Get("test:name")
	if err != nil {
		fmt.Printf("error : %v\n", err)
		return
	}
	fmt.Println("s:" + s)

	//List 操作
	// 从列表的尾部进行添加
	err = Cache.RPush("message", "01")
	if err != nil {
		fmt.Printf("rpush message error: %v\n", err)
		return
	}
	// 从列表的头部进行添加
	err = Cache.LPush("message", "02")
	if err != nil {
		fmt.Printf("lpush message error: %v\n", err)
		return
	}
	// 从列表尾部弹出元素
	message, err := Cache.RPop("message")
	if err != nil {
		fmt.Printf("rpop message error: %v\n", err)
		return
	}
	fmt.Println("rpop message: ", message)
	// 获取长度
	length, err := Cache.LLen("message")
	if err != nil {
		fmt.Printf("llen message error: %v\n", err)
		return
	}
	fmt.Println("message length: ", length)

	//Hash 操作
	key := "string:hash"
	Cache.HSet(key, "name", "张三")
	Cache.HSet(key, "phone", "18234554345")
	Cache.HSet(key, "age", "28")
	//获取全部hash对象
	all, _ := Cache.HGetAll(key)
	fmt.Println(all)
	//修改已存在的字段
	Cache.HSet(key, "name", "李四")
	//获取指定字段
	name, _ := Cache.HGet(key, "name")
	fmt.Println(name)
	existsName, _ := Cache.HExists(key, "name")
	existsId, _ := Cache.HExists(key, "id")
	fmt.Printf("name 字段是否存在 %v\n", existsName)
	fmt.Printf("id 字段是否存在 %v\n", existsId)
	Cache.HDel(key, "name")
	existsName, _ = Cache.HExists(key, "name")
	fmt.Printf("name 字段是否存在 %v\n", existsName)
	getAll, _ := Cache.HGetAll(key)
	fmt.Println(getAll)

	//SET 操作
	key1 := "string:set"
	Cache.SAdd(key1, "phone")
	err2 := Cache.SAdd(key1, "hahh")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	//获取全部hash对象
	all1, _ := Cache.SCard(key1)
	fmt.Println(all1)
	members, err2 := Cache.SMembers(key1)
	fmt.Println(members)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}
