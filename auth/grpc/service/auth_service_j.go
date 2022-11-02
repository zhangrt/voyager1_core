package service

import (
	"sync"

	luna "github.com/zhangrt/voyager1_core/auth/luna"
	pb "github.com/zhangrt/voyager1_core/com/gs/voyager1_core/auth/grpc/pb"
	"github.com/zhangrt/voyager1_core/util"

	"context"
)

// 懒加载初始化
var (
	casbinJ luna.Casbin
	onceJ   sync.Once
)

// Grpc Server 接口实现
// AuthServiceJ 接口实现，需要依赖注入luna.JWT luna.Casbin
// JWT、Casbin的接口实现在GRPC服务启动前注入
type AuthServiceJ struct {
	// grpc 必须继承这个proto接口 并进行实现
	pb.UnimplementedAuthServiceServer
}

// 通过Casbin校验角色权限 - authorityId、path、method
func (auth *AuthServiceJ) GrantedAuthority(c context.Context, p *pb.Policy) (*pb.Result, error) {
	onceJ.Do(func() {
		casbinJ = luna.NewCasbin()
	})
	result := new(pb.Result)
	var err_ error
	s, m, claims, n, err := luna.ReadAuthentication(p.Token)
	result.Success = s
	result.Msg = m
	result.Claims = util.GrpcLunaClaimsTransformProtoClaimsJ(claims)
	result.NewToken = n
	// token不合法或过期等情况
	if !s {
		result.Msg = m
		return result, err
	}
	if p.Path == "" || p.Method == "" {
		result.Success = false
		if n == "" {
			result.Msg = "Unkonw policy"
		}
		return result, err
	}

	// 校验角色信息权限，只要有一个角色有权限即通过
	e := casbinJ.Casbin()
	for i, roleId := range claims.RoleIds {
		success, err := e.Enforce(roleId, p.Path, p.Method)
		if err != nil {
			err_ = err
			result.Msg = err.Error()
			// 错误
			break
		}
		if success {
			result.Success = success
			err_ = nil
			break
		}
		if !success && (i == len(claims.RoleIds)-1) {
			result.Success = false
			result.Msg = "insufficient privileges"
			err_ = nil
		}
	}
	return result, err_

}

// 通过Token获取用户信息
func (auth *AuthServiceJ) GetUser(c context.Context, p *pb.Authentication) (*pb.User, error) {
	user := new(pb.User)
	claims, _ := luna.GetUser(p.Token)
	if claims != nil {
		user.Claims = util.GrpcLunaClaimsTransformProtoClaimsJ(claims)
		user.ID = claims.Id
		user.RoleIds = claims.RoleIds
	}
	return user, nil
}
