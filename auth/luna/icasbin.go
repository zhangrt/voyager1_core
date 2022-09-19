package luna

import (
	"github.com/casbin/casbin/v2"
)

// 权限接口
type Casbin interface {
	// 更新casbin权限
	UpdateCasbin(authorityId string, casbinInfos []CasbinInfo) error
	// 获取权限列表
	GetPolicyPathByAuthorityId(authorityId string) (pathMaps []CasbinInfo)
	// 清除匹配的权限
	ClearCasbin(v int, p ...string) bool
	// 持久化到数据库  引入自定义规则
	Casbin() *casbin.SyncedEnforcer
}

// @author: [zhoujiajun](zhoujiajun@gsafety.com)
// 权限实现
func NewCasbin() Casbin {

	return &CasbinService{}
}
