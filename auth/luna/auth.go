package luna

import (
	"github.com/gin-gonic/gin"
)

func Enforce(c *gin.Context) (bool, error) {
	waitUse, _ := GetClaims(c)
	// 获取请求的PATH
	obj := c.Request.URL.Path
	// 获取请求方法
	act := c.Request.Method
	// 获取用户的角色
	sub := waitUse.AuthorityId

	return CheckPolicy(sub, obj, act)
}

func CheckPolicy(sub string, obj string, act string) (bool, error) {
	auth := NewCasbin()
	e := auth.Casbin()
	// 判断策略中是否存在
	success, err := e.Enforce(sub, obj, act)
	return success, err
}
