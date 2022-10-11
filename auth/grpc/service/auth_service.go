package service

import (
	"fmt"
	"sync"
	"time"

	pb "github.com/zhangrt/voyager1_core/auth/grpc/pb"
	luna "github.com/zhangrt/voyager1_core/auth/luna"
	"github.com/zhangrt/voyager1_core/constant"
	"github.com/zhangrt/voyager1_core/global"
	"github.com/zhangrt/voyager1_core/util"
	"go.uber.org/zap"

	"context"
)

// 懒加载初始化
var (
	jwt    luna.JWT
	casbin luna.Casbin
	once   sync.Once
)

// Grpc Server 接口实现
// AuthService 接口实现，需要依赖注入luna.JWT luna.Casbin
// JWT、Casbin的接口实现在GRPC服务启动前注入
type AuthService struct {
	// grpc 必须继承这个proto接口 并进行实现
	pb.UnimplementedAuthServiceServer
}

// 读取Token验证Token合法性与过期时间校验
func (auth *AuthService) ReadAuthentication(c context.Context, p *pb.Token) (*pb.Result, error) {
	once.Do(func() {
		jwt = luna.NewJWT()
	})
	result := new(pb.Result)
	token := p.Token
	if token == "" {
		result.Msg = "Not logged in or accessed illegally"
		result.Success = false
		return result, nil
	}
	if jwt.IsBlacklist(token) {
		result.Msg = "Your account is off-site logged in or the token is invalid"
		result.Success = false
		return result, nil
	}
	j := luna.NewTOKEN()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		if err == luna.TokenExpired {
			result.Msg = "Authorization has expired"
			result.Success = false
			return result, nil
		}
		result.Msg = err.Error()
		result.Success = false
		return result, err
	}
	// 解析token成功
	result.Success = true
	result.Claims = util.GrpcLunaClaimsTransformProtoClaims(claims)

	// 判断是否需要生成Newtoken
	now := time.Now().Unix()
	if claims.ExpiresAt-now < claims.BufferTime {
		claims.ExpiresAt = now + global.G_CONFIG.JWT.ExpiresTime
		newToken, _ := j.CreateTokenByOldToken(token, *claims)
		newClaims, _ := j.ParseToken(newToken)
		// 将 New Token 存在 Msg 中
		result.Msg = fmt.Sprintf("%s"+constant.MARKER+"%d", newToken, newClaims.ExpiresAt)
		// 存 New Claims
		result.Claims = util.GrpcLunaClaimsTransformProtoClaims(newClaims)
		// 单点, 在 Server 端进行拉黑
		if !global.G_CONFIG.System.UseMultipoint {
			RedisJwtToken, err := jwt.GetCacheJWT(newClaims.Account)
			if err != nil {
				global.G_LOG.Error("get redis jwt failed", zap.Error(err))
			} else { // 当之前的取成功时才进行拉黑操作
				_ = jwt.JsonInBlacklist(luna.JwtBlacklist{Jwt: RedisJwtToken})
			}
			// 无论如何都要记录当前的活跃状态
			_ = jwt.SetCacheJWT(newToken, newClaims.Account)
		}
	}
	return result, nil
}

// 通过Casbin校验角色权限 - authorityId、path、method
func (auth *AuthService) GrantedAuthority(c context.Context, p *pb.Policy) (*pb.Result, error) {
	once.Do(func() {
		casbin = luna.NewCasbin()
	})
	result := new(pb.Result)
	e := casbin.Casbin()
	success, err := e.Enforce(p.AuthorityId, p.Path, p.Method)
	result.Success = success
	return result, err

}

// 通过Token获取用户信息
func (auth *AuthService) GetUser(c context.Context, p *pb.Token) (*pb.User, error) {
	user := new(pb.User)
	claims, _ := luna.GetUser(p.Token)
	if claims != nil {
		user.Claims = util.GrpcLunaClaimsTransformProtoClaims(claims)
		user.UserID = int64(claims.ID)
		user.UUID = claims.UUID.String()
		user.AuthorityId = claims.RoleId
	}
	return user, nil
}
