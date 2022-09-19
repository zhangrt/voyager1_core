package global

import (
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/xyy277/gallery/config"
)

var (
	G_DB                  *gorm.DB
	GS_DBList             map[string]*gorm.DB
	G_REDIS_STANDALONE    *redis.Client
	G_REDIS_CLUSTER       *redis.ClusterClient
	G_REDIS_CLUSTER_MOD   bool
	G_CONFIG              config.Server
	G_VP                  *viper.Viper
	G_LOG                 *zap.Logger
	G_Concurrency_Control = &singleflight.Group{}
	BlackCache            local_cache.Cache
	lock                  sync.RWMutex
)
