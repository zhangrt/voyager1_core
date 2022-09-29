package luna

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

// 权限接口 在这里不提供实现接口，由调用者实现
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

var c Casbin

// @author: [zhoujiajun](zhoujiajun@gsafety.com)
// 权限实现
func NewCasbin() Casbin {
	if c != nil {
		return c
	}
	fmt.Errorf("Interface Casbin not implemented")
	return c
}

type UnimplementedCasbin struct {
}

// 更新casbin权限
func (UnimplementedCasbin) UpdateCasbin(authorityId string, casbinInfos []CasbinInfo) error {
	return fmt.Errorf("method UpdateCasbin not implemented")
}

// 获取权限列表
func (UnimplementedCasbin) GetPolicyPathByAuthorityId(authorityId string) (pathMaps []CasbinInfo) {
	return nil
}

// 清除匹配的权限
func (UnimplementedCasbin) ClearCasbin(v int, p ...string) bool {
	return false
}

// 持久化到数据库  引入自定义规则
func (UnimplementedCasbin) Casbin() *casbin.SyncedEnforcer {
	return nil
}

func RegisterCasbin(casb Casbin) {
	c = casb
}
