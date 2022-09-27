package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/zhangrt/voyager1_core/config"
)

var (
	// core 去除 gorm相关代码
	// G_DB                  *gorm.DB
	// GS_DBList             map[string]*gorm.DB
	G_REDIS_STANDALONE    *redis.Client
	G_REDIS_CLUSTER       *redis.ClusterClient
	G_REDIS_CLUSTER_MOD   bool
	G_CONFIG              config.Server
	G_VP                  *viper.Viper
	G_LOG                 *zap.Logger
	G_Concurrency_Control = &singleflight.Group{}
	// BlackCache            local_cache.Cache
	// lock sync.RWMutex
)
