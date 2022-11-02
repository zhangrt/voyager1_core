package luna

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zhangrt/voyager1_core/global"
	"go.uber.org/zap"
)

var (
	auth Casbin
	once sync.Once
	ijwt JWT
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

// token鉴权 & 校验path method的权限
func CheckAuth(token string, obj string, act string) (bool, string, string, error) {
	s, m, c, n, e := ReadAuthentication(token)
	if !s {
		return s, m, n, e
	}
	s, e = CheckPolicy(c.RoleIds, obj, act)
	if e != nil {
		return s, e.Error(), n, e
	}
	return s, m, n, e
}

func CheckPolicy(subs []string, obj string, act string) (bool, error) {

	e := auth.Casbin()
	// 判断策略中是否存在
	for _, sub := range subs {
		success, err := e.Enforce(sub, obj, act)
		if err != nil {
			return false, err
		}
		if success {
			return true, nil
		}
	}
	return false, nil
}

// 读取Token验证Token合法性与过期时间校验
func ReadAuthentication(token string) (bool, string, *CustomClaims, string, error) {
	once.Do(func() {
		ijwt = NewJWT()
		auth = NewCasbin()
	})
	var success = false
	var msg = ""
	var newToken = ""
	if token == "" {
		msg = "Not logged in or accessed illegally"
		success = false
		return success, msg, nil, newToken, nil
	}
	if ijwt.IsBlacklist(token) {
		msg = "Your account is off-site logged in or the token is invalid"
		success = false
		return success, msg, nil, newToken, nil
	}
	j := NewTOKEN()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		if err == TokenExpired {
			msg = "Authorization has expired"
			success = false
			return success, msg, nil, newToken, nil
		}
		msg = err.Error()
		success = false
		return success, msg, nil, newToken, err
	}
	// 解析token成功
	success = true

	// 判断是否需要生成Newtoken
	now := time.Now().Unix()
	if claims.ExpiresAt-now < claims.BufferTime {
		claims.ExpiresAt = now + global.G_CONFIG.JWT.ExpiresTime
		newToken, _ = j.CreateTokenByOldToken(token, *claims)
		newClaims, _ := j.ParseToken(newToken)
		msg = "The token is about to expire, so we create a new token"
		// 存 New Claims
		claims = newClaims
		// 单点, 在 Server 端进行拉黑
		if !global.G_CONFIG.System.UseMultipoint {
			RedisJwtToken, err := ijwt.GetCacheJWT(newClaims.Account)
			if err != nil {
				global.G_LOG.Error("get redis jwt failed", zap.Error(err))
			} else { // 当之前的取成功时才进行拉黑操作
				_ = ijwt.JsonInBlacklist(JwtBlacklist{Jwt: RedisJwtToken})
			}
			// 无论如何都要记录当前的活跃状态
			_ = ijwt.SetCacheJWT(newToken, newClaims.Account)
		}
	}
	return success, msg, claims, newToken, nil
}
