package handler

import (
	"github.com/zhangrt/voyager1_core/auth/luna"
	"github.com/zhangrt/voyager1_core/auth/star"
	"github.com/zhangrt/voyager1_core/global"
	"github.com/zhangrt/voyager1_core/global/response"

	"github.com/gin-gonic/gin"
)

// 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取jwt claims信息
		claims, e := c.Get(global.G_CONFIG.AUTHKey.User)
		if e {
			// 获取请求人角色ID
			authorityId := claims.(*luna.CustomClaims).AuthorityId
			// 获取请求的PATH
			obj := c.Request.URL.Path
			// 获取请求方法
			act := c.Request.Method
			success := star.NewAUTH().GrantedAuthority(authorityId, obj, act)
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
