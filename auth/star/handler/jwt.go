package handler

import (
	auth "github.com/xyy277/gallery/auth/luna"
	"github.com/xyy277/gallery/auth/star"
	"github.com/xyy277/gallery/global/response"

	"github.com/xyy277/gallery/global"

	"github.com/gin-gonic/gin"
)

var jwt = auth.NewJWT()

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get(global.G_CONFIG.AUTHKey.Token)
		t, m, claims := star.NewAUTH().RemoteAuthentication(token)
		if !t {
			response.FailWithDetailed(gin.H{"reload": true}, m, c)
			c.Abort()
		} else {
			// set claims
			c.Set(global.G_CONFIG.AUTHKey.User, claims)
			c.Next()
		}
	}
}
