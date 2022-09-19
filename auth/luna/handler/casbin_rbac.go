package handler

import (
	"github.com/zhangrt/voyager1_core/global/response"

	auth "github.com/zhangrt/voyager1_core/auth/luna"

	"github.com/gin-gonic/gin"
)

// 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		success, _ := auth.Enforce(c)
		if success {
			c.Next()
		} else {
			response.FailWithDetailed(gin.H{}, "insufficient privileges", c)
			c.Abort()
			return
		}
	}
}
