package handler

import (
	"strconv"
	"strings"
	"sync"

	"github.com/zhangrt/voyager1_core/auth/star"
	"github.com/zhangrt/voyager1_core/constant"
	"github.com/zhangrt/voyager1_core/global/response"

	"github.com/zhangrt/voyager1_core/global"

	"github.com/gin-gonic/gin"
)

var (
	auth star.AUTH
	once sync.Once
)

// 拦截器JWT 拦截请求，GRPC方式校验Token合法性及权限验证
// JWT拦截器 传入impl选择不同通信方式的接口实现
// star服务可使用此默认handler，也可依据特殊业务通过提供的Auth接口自行实现
func JWTAuth(impl string) gin.HandlerFunc {
	once.Do(func() {
		auth = star.NewAUTH(impl)
	})
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get(global.G_CONFIG.AUTHKey.Token)
		// 获取请求的PATH
		path := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		t, m, claims := auth.GrantedAuthority(token, path, act)
		if !t {
			response.FailWithDetailed(gin.H{"reload": true}, m, c)
			c.Abort()
		} else {
			// refresh token
			s := strings.Split(m, constant.MARKER)
			if len(s) > 1 {
				c.Header(global.G_CONFIG.AUTHKey.RefreshToken, s[0])
				c.Header(global.G_CONFIG.AUTHKey.RefreshExpiresAt, strconv.FormatInt(claims.ExpiresAt, 10))
			}
			// set claims
			c.Set(global.G_CONFIG.AUTHKey.User, claims)
			c.Next()
		}
	}
}
