package handler

import (
	"github.com/zhangrt/voyager1_core/auth/luna"
	"github.com/zhangrt/voyager1_core/global"
	"github.com/zhangrt/voyager1_core/global/response"

	"github.com/gin-gonic/gin"
)

// Casbin 拦截器  传入impl选择不同通信方式的接口实现
// 集成star的star服务可以使用此默认Handler也可以根据特定业务逻辑调用提供的Auth接口自己实现
func CasbinHandler(impl string) gin.HandlerFunc {

	return func(c *gin.Context) {
		// 获取jwt claims信息
		claims, e := c.Get(global.G_CONFIG.AUTHKey.User)
		if e {
			// 获取请求人角色ID
			authorityId := claims.(*luna.CustomClaims).RoleId
			// 获取请求的PATH
			obj := c.Request.URL.Path
			// 获取请求方法
			act := c.Request.Method
			success := auth.GrantedAuthority(authorityId, obj, act)
			if success {
				c.Next()
			} else {
				response.FailWithDetailed(gin.H{}, "insufficient privileges", c)
				c.Abort()
				return
			}
		} else {
			response.FailWithDetailed(gin.H{}, "No role information was obtained ", c)
			c.Abort()
		}

	}
}
