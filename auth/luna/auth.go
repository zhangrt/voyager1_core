package luna

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	auth Casbin
	once sync.Once
)

// 校验用户角色的Policy

func Enforce(c *gin.Context) (bool, error) {
	waitUse, _ := GetClaims(c)
	// 获取请求的PATH
	obj := c.Request.URL.Path
	// 获取请求方法
	act := c.Request.Method
	// 获取用户的角色
	sub := waitUse.RoleIds

	return CheckPolicy(sub, obj, act)
}

func CheckPolicy(sub []string, obj string, act string) (bool, error) {
	once.Do(func() {
		auth = NewCasbin()
	})
	e := auth.Casbin()
	// 判断策略中是否存在
	success, err := e.Enforce(sub, obj, act)
	return success, err
}
