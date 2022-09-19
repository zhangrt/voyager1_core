package cache

import (
	"context"
	"time"

	v8 "github.com/go-redis/redis/v8"

	"github.com/xyy277/gallery/global"
)

type Result struct {
	val int64
}

// check nil
func Nil() bool {
	if global.G_REDIS_CLUSTER_MOD {
		return global.G_REDIS_CLUSTER == nil
	} else {
		return global.G_REDIS_STANDALONE == nil
	}
}

func ExistsResultByKey(key string) (int64, error) {
	var r int64
	var err error
	if global.G_REDIS_CLUSTER_MOD {
		r, err = global.G_REDIS_CLUSTER.Exists(context.Background(), key).Result()
	} else {
		r, err = global.G_REDIS_STANDALONE.Exists(context.Background(), key).Result()

	}
	return r, err
}

func GetInt(key string) (int, error) {
	var r int
	var err error
	if global.G_REDIS_CLUSTER_MOD {
		r, err = global.G_REDIS_CLUSTER.Get(context.Background(), key).Int()
	} else {
		r, err = global.G_REDIS_STANDALONE.Get(context.Background(), key).Int()

	}
	return r, err
}

func Incr(key string) error {
	var err error
	if global.G_REDIS_CLUSTER_MOD {
		err = global.G_REDIS_CLUSTER.Incr(context.Background(), key).Err()
	} else {
		err = global.G_REDIS_STANDALONE.Incr(context.Background(), key).Err()

	}
	return err
}

func TxPipeline() v8.Pipeliner {
	var pipe v8.Pipeliner
	if global.G_REDIS_CLUSTER_MOD {
		pipe = global.G_REDIS_CLUSTER.TxPipeline()
	} else {
		pipe = global.G_REDIS_STANDALONE.TxPipeline()

	}
	return pipe
}

func PTTL(key string) (time.Duration, error) {
	var t time.Duration
	var err error
	if global.G_REDIS_CLUSTER_MOD {
		t, err = global.G_REDIS_CLUSTER.PTTL(context.Background(), key).Result()
	} else {
		t, err = global.G_REDIS_STANDALONE.PTTL(context.Background(), key).Result()
	}
	return t, err
}

func GetResult(key string) (string, error) {
	var r string
	var err error
	if global.G_REDIS_CLUSTER_MOD {
		r, err = global.G_REDIS_CLUSTER.Get(context.Background(), key).Result()
	} else {
		r, err = global.G_REDIS_STANDALONE.Get(context.Background(), key).Result()

	}
	return r, err
}

func Set(key string, value interface{}, expiration time.Duration) error {
	var err error
	if global.G_REDIS_CLUSTER_MOD {
		err = global.G_REDIS_CLUSTER.Set(context.Background(), key, value, expiration).Err()
	} else {
		err = global.G_REDIS_STANDALONE.Set(context.Background(), key, value, expiration).Err()
	}
	return err
}
