package handler

import (
	"strconv"
	"time"

	auth "github.com/zhangrt/voyager1_core/auth/luna"

	"github.com/zhangrt/voyager1_core/global"
	"github.com/zhangrt/voyager1_core/global/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var jwt = auth.NewJWT()

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get(global.G_CONFIG.AUTHKey.Token)
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "Not logged in or accessed illegally", c)
			c.Abort()
			return
		}
		if jwt.IsBlacklist(token) {
			response.FailWithDetailed(gin.H{"reload": true}, "Your account is off-site logged in or the token is invalid", c)
			c.Abort()
			return
		}
		j := auth.NewTOKEN()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == auth.TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "Authorization has expired", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开
		//if err, _ = userService.FindUserByUuid(claims.UUID.String()); err != nil {
		//	_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
		//	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		//	c.Abort()
		//}
		now := time.Now().Unix()
		if claims.ExpiresAt-now < claims.BufferTime {
			claims.ExpiresAt = now + global.G_CONFIG.JWT.ExpiresTime
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header(global.G_CONFIG.AUTHKey.RefreshToken, newToken)
			c.Header(global.G_CONFIG.AUTHKey.RefreshExpiresAt, strconv.FormatInt(newClaims.ExpiresAt, 10))
			// 单点
			if !global.G_CONFIG.System.UseMultipoint {
				RedisJwtToken, err := jwt.GetCacheJWT(newClaims.Account)
				if err != nil {
					global.G_LOG.Error("get redis jwt failed", zap.Error(err))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = jwt.JsonInBlacklist(auth.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwt.SetCacheJWT(newToken, newClaims.Account)
			}
		}
		// set claims
		c.Set(global.G_CONFIG.AUTHKey.User, claims)
		c.Next()
	}
}
