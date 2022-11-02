package handler

import (
	"strconv"

	auth "github.com/zhangrt/voyager1_core/auth/luna"

	"github.com/zhangrt/voyager1_core/global"
	"github.com/zhangrt/voyager1_core/global/response"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get(global.G_CONFIG.AUTHKey.Token)
		success, msg, claims, newToken, err := auth.ReadAuthentication(token)
		if success {
			// refresh token
			if newToken != "" {
				c.Header(global.G_CONFIG.AUTHKey.RefreshToken, newToken)
				c.Header(global.G_CONFIG.AUTHKey.RefreshExpiresAt, strconv.FormatInt(claims.ExpiresAt, 10))
			}
			// set claims
			c.Set(global.G_CONFIG.AUTHKey.User, claims)
			c.Next()
		} else {
			if err != nil {
				msg = err.Error()
			}
			response.FailWithDetailed(gin.H{global.G_CONFIG.AUTHKey.Reload: true}, msg, c)
			c.Abort()
		}

	}
}
