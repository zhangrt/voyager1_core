package auth

import (
	"fmt"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/xyy277/gallery/config"
	"github.com/xyy277/gallery/global"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	InitAuthKey bool
)

type Init struct{}

func NewInit() *Init {
	fmt.Printf(`
	_____/\\\\\\\\\\\\_______/\\\\\\\\\\\\__
 ___/\\\//////////_____/\\\//////////__
  __/\\\_______________/\\\_____________
   _\/\\\____/\\\\\\\__\/\\\____/\\\\\\\_
    _\/\\\___\/////\\\__\/\\\___\/////\\\_
     _\/\\\_______\/\\\__\/\\\_______\/\\\_
      _\/\\\_______\/\\\__\/\\\_______\/\\\_ 
       _\//\\\\\\\\\\\\/___\//\\\\\\\\\\\\/__
        __\////////////______\////////////____
	`)
	return &Init{}
}

func (init *Init) DB(db *gorm.DB) *Init {
	global.G_DB = db
	return init
}

func (init *Init) SetRedisMod(b bool) *Init {
	global.G_REDIS_CLUSTER_MOD = b
	return init
}

func (init *Init) RedisStandalone(r *redis.Client) *Init {
	global.G_REDIS_CLUSTER_MOD = false
	global.G_REDIS_STANDALONE = r
	return init
}

func (init *Init) RedisCluster(rs *redis.ClusterClient) *Init {
	global.G_REDIS_CLUSTER_MOD = true
	global.G_REDIS_CLUSTER = rs
	return init
}

func (init *Init) Viper(vp *viper.Viper) *Init {
	global.G_VP = vp
	return init
}

func (init *Init) Zap(zap *zap.Logger) *Init {
	global.G_LOG = zap
	return init
}

func (init *Init) BlackCache(cache local_cache.Cache) *Init {
	global.BlackCache = cache
	return init
}

func (init *Init) Config(config config.Server) *Init {
	global.G_CONFIG = config
	return init
}

func (init *Init) ConSystem(c config.System) *Init {
	global.G_CONFIG.System = c
	return init
}

func (init *Init) ConfigAuth(c config.AUTHKey) *Init {
	InitAuthKey = true
	global.G_CONFIG.AUTHKey = c
	return init
}

func (init *Init) ConfigCasbin(c config.Casbin) *Init {
	global.G_CONFIG.Casbin = c
	return init
}

func (init *Init) ConfigJwt(c config.JWT) *Init {
	global.G_CONFIG.JWT = c
	return init
}

func (init *Init) ConfigMinio(c config.Minio) *Init {
	global.G_CONFIG.Minio = c
	return init
}

func (init *Init) ConfigZinx(c config.Zinx) *Init {
	global.G_CONFIG.Zinx = c
	return init
}
