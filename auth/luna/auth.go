package luna

import (
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zhangrt/voyager1_core/constant"
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
func CheckAuth(token string, obj string, act string) (bool, string, error) {
	s, m, c, e := ReadAuthentication(token)
	if !s {
		return s, m, e
	}
	s, e = CheckPolicy(c.RoleIds, obj, act)
	if e != nil {
		return s, e.Error(), e
	}
	return s, m, e
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

// 读取Token验证Token合法性与过期时间校验
func ReadAuthentication(token string) (bool, string, *CustomClaims, error) {
	var success = false
	var msg = ""
	once.Do(func() {
		ijwt = NewJWT()
	})
	if token == "" {
		msg = "Not logged in or accessed illegally"
		success = false
		return success, msg, nil, nil
	}
	if ijwt.IsBlacklist(token) {
		msg = "Your account is off-site logged in or the token is invalid"
		success = false
		return success, msg, nil, nil
	}
	j := NewTOKEN()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		if err == TokenExpired {
			msg = "Authorization has expired"
			success = false
			return success, msg, nil, nil
		}
		msg = err.Error()
		success = false
		return success, msg, nil, err
	}
	// 解析token成功
	success = true

	// 判断是否需要生成Newtoken
	now := time.Now().Unix()
	if claims.ExpiresAt-now < claims.BufferTime {
		claims.ExpiresAt = now + global.G_CONFIG.JWT.ExpiresTime
		newToken, _ := j.CreateTokenByOldToken(token, *claims)
		newClaims, _ := j.ParseToken(newToken)
		// 将 New Token 存在 Msg 中
		msg = fmt.Sprintf("%s"+constant.MARKER+"%d", newToken, newClaims.ExpiresAt)
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
	return success, msg, claims, nil
}
