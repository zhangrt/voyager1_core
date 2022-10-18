package cache

import (
	"runtime"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/zhangrt/voyager1_core/global"
)

//实现具体的Cache：RedisClusterCache
type RedisClusterCache struct {
}

var redisCluster *redis.ClusterClient

// func getRedisCluster() *redis.ClusterClient {
// 	if redisCluster == nil {
// 		lock.Lock()
// 		defer lock.Unlock()
// 		if redisCluster == nil {
// 			redisCluster = createRedisCluster()
// 		}
// 	}
// 	return redisCluster
// }

func createRedisCluster() (*redis.ClusterClient, error) {
	redisCluster := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    global.G_CONFIG.Cache.Addrs,
		Username: global.G_CONFIG.Cache.Username,
		Password: global.G_CONFIG.Cache.Password,
		ReadOnly: false,

		//每一个redis.Client的连接池容量及闲置连接数量，而不是clusterClient总体的连接池大小。
		//实际上没有总的连接池而是由各个redis.Client自行去实现和维护各自的连接池。
		PoolSize:     20 * runtime.NumCPU(), // 连接池最大socket连接数，默认为5倍CPU数， 5 * runtime.NumCPU
		MinIdleConns: 10,                    //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量。

		//命令执行失败时的重试策略
		MaxRetries:      0,                      // 命令执行失败时，最多重试多少次，默认为0即不重试
		MinRetryBackoff: 8 * time.Millisecond,   //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
		MaxRetryBackoff: 512 * time.Millisecond, //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔

		//超时
		DialTimeout:  5 * time.Second, //连接建立超时时间，默认5秒。
		ReadTimeout:  3 * time.Second, //读超时，默认3秒， -1表示取消读超时
		WriteTimeout: 3 * time.Second, //写超时，默认等于读超时，-1表示取消读超时
		PoolTimeout:  4 * time.Second, //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。

		IdleTimeout: 5 * time.Minute, //闲置超时，默认5分钟，-1表示取消闲置超时检查
		MaxConnAge:  0 * time.Second, //连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接
	})
	_, err := redisCluster.Ping(ctx).Result()
	return redisCluster, err
}

func (redisClusterCache *RedisClusterCache) InitCache() error {
	var err error
	if redisCluster == nil {
		lock.Lock()
		defer lock.Unlock()
		if redisCluster == nil {
			redisCluster, err = createRedisCluster()
		}
	}
	return err
}

func (redisClusterCache *RedisClusterCache) Get(key string) (string, error) {
	result, err := redisCluster.Get(ctx, key).Result()
	return result, err
}

func (redisClusterCache *RedisClusterCache) Set(key string, value interface{}) error {
	err := redisCluster.Set(ctx, key, value, time.Second*0).Err()
	return err
}

func (redisClusterCache *RedisClusterCache) SetX(key string, value interface{}, expiration time.Duration) error {
	err := redisCluster.Set(ctx, key, value, expiration).Err()
	return err
}

func (redisClusterCache *RedisClusterCache) LPush(key string, value ...interface{}) error {
	err := redisCluster.LPush(ctx, key, value).Err()
	return err
}

func (redisClusterCache *RedisClusterCache) RPush(key string, value ...interface{}) error {
	err := redisCluster.RPush(ctx, key, value).Err()
	return err
}

func (redisClusterCache *RedisClusterCache) LPop(key string) (interface{}, error) {
	result, err := redisCluster.LPop(ctx, key).Result()
	return result, err
}

func (redisClusterCache *RedisClusterCache) RPop(key string) (interface{}, error) {
	result, err := redisCluster.RPop(ctx, key).Result()
	return result, err
}

func (redisClusterCache *RedisClusterCache) BLPop(timeout time.Duration, key string) (interface{}, error) {
	result, err := redisCluster.BLPop(ctx, timeout, key).Result()
	return result, err
}

func (redisClusterCache *RedisClusterCache) BRPop(timeout time.Duration, key string) (interface{}, error) {
	result, err := redisCluster.BRPop(ctx, timeout, key).Result()
	return result, err
}

func (redisClusterCache *RedisClusterCache) LLen(key string) (int64, error) {
	result, err := redisCluster.LLen(ctx, key).Result()
	return result, err
}

func (redisClusterCache *RedisClusterCache) LRange(key string, start, end int64) ([]string, error) {
	result, err := redisCluster.LRange(ctx, key, start, end).Result()
	return result, err
}

func (redisClusterCache *RedisClusterCache) HSet(hashKey, key string, value interface{}) error {
	err := redisCluster.HSet(ctx, hashKey, key, value).Err()
	return err
}

func (redisClusterCache *RedisClusterCache) HGet(hashKey, key string) (interface{}, error) {
	result, err := redisCluster.HGet(ctx, hashKey, key).Result()
	return result, err
}

func (redisClusterCache *RedisClusterCache) HGetAll(hashKey string) (map[string]string, error) {
	result, err := redisCluster.HGetAll(ctx, hashKey).Result()
	return result, err
}

func (redisClusterCache *RedisClusterCache) HDel(hashKey string, key ...string) error {
	err := redisCluster.HDel(ctx, hashKey, key...).Err()
	return err
}

func (redisClusterCache *RedisClusterCache) HExists(hashKey, key string) (bool, error) {
	result, err := redisCluster.HExists(ctx, hashKey, key).Result()
	return result, err
}

func (redisClusterCache *RedisClusterCache) SAdd(key string, values ...interface{}) error {
	err := redisCluster.SAdd(ctx, key, values).Err()
	return err
}

func (redisClusterCache *RedisClusterCache) SCard(key string) (int64, error) {
	result, err := redisCluster.SCard(ctx, key).Result()
	return result, err
}

func (redisClusterCache *RedisClusterCache) SMembers(key string) ([]string, error) {
	result, err := redisCluster.SMembers(ctx, key).Result()
	return result, err
}

func (redisClusterCache *RedisClusterCache) SRem(key string, value interface{}) error {
	err := redisCluster.SRem(ctx, key, value).Err()
	return err
}

func (redisClusterCache *RedisClusterCache) SPop(key string) (interface{}, error) {
	result, err := redisCluster.SPop(ctx, key).Result()
	return result, err
}

func (redisClusterCache *RedisClusterCache) Expire(key string, duration time.Duration) error {
	err := redisCluster.Expire(ctx, key, duration).Err()
	return err
}

func (redisClusterCache *RedisClusterCache) ExpireAt(key string, duration time.Time) error {
	err := redisCluster.ExpireAt(ctx, key, duration).Err()
	return err
}

func (redisClusterCache *RedisClusterCache) TTL(key string) (time.Duration, error) {
	result, err := redisCluster.TTL(ctx, key).Result()
	return result, err
}
