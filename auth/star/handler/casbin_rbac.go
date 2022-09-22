package handler

import (
	"github.com/zhangrt/voyager1_core/auth/star"
	"github.com/zhangrt/voyager1_core/global/response"

	"github.com/gin-gonic/gin"
)

// 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		success := star.NewAUTH().GrantedAuthority(obj, act)
		if success {
			c.Next()
		} else {
			response.FailWithDetailed(gin.H{}, "insufficient privileges", c)
			c.Abort()
			return
		}
	}
}
