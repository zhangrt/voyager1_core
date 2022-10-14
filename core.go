package voyager1_core

import (
	"fmt"

	"github.com/zhangrt/voyager1_core/config"
	"github.com/zhangrt/voyager1_core/global"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	InitAuthKey bool
)

// 初始化配置信息
type Core struct{}

func New() *Core {
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
	fmt.Println()
	return &Core{}
}

// func (Core *Core) DB(db *gorm.DB) *Core {
// 	global.G_DB = db
// 	return Core
// }

func (Core *Core) SetRedisMod(b bool) *Core {
	global.G_REDIS_CLUSTER_MOD = b
	return Core
}

func (Core *Core) RedisStandalone(r *redis.Client) *Core {
	global.G_REDIS_CLUSTER_MOD = false
	global.G_REDIS_STANDALONE = r
	return Core
}

func (Core *Core) RedisCluster(rs *redis.ClusterClient) *Core {
	global.G_REDIS_CLUSTER_MOD = true
	global.G_REDIS_CLUSTER = rs
	return Core
}

func (Core *Core) Viper(vp *viper.Viper) *Core {
	global.G_VP = vp
	return Core
}

func (Core *Core) Zap(zap *zap.Logger) *Core {
	global.G_LOG = zap
	return Core
}

// func (Core *Core) BlackCache(cache local_cache.Cache) *Core {
// 	global.BlackCache = cache
// 	return Core
// }

func (Core *Core) Config(config config.Server) *Core {
	global.G_CONFIG = config
	return Core
}

func (Core *Core) ConSystem(c config.System) *Core {
	global.G_CONFIG.System = c
	return Core
}

func (Core *Core) ConfigAuth(c config.AUTHKey) *Core {
	InitAuthKey = true
	global.G_CONFIG.AUTHKey = c
	return Core
}

func (Core *Core) ConfigCasbin(c config.Casbin) *Core {
	global.G_CONFIG.Casbin = c
	return Core
}

func (Core *Core) ConfigJwt(c config.JWT) *Core {
	global.G_CONFIG.JWT = c
	return Core
}

func (Core *Core) ConfigMinio(c config.Minio) *Core {
	global.G_CONFIG.Minio = c
	return Core
}

func (Core *Core) ConfigZinx(c config.Zinx) *Core {
	global.G_CONFIG.Zinx = c
	return Core
}

func (Core *Core) ConfigGrpc(c config.Grpc) *Core {
	global.G_CONFIG.Grpc = c
	return Core
}
