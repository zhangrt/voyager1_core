package star

import (
	"github.com/zhangrt/voyager1_core/auth/luna"
	"github.com/zhangrt/voyager1_core/constant"
)

// star 请求 luna的接口进行鉴权与授权
type AUTH interface {
	// 通过token请求鉴权
	ReadAuthentication(token string) (bool, string, *luna.CustomClaims)
	// 通过request中的path和method请求验证策略
	GrantedAuthority(authorityId string, path string, method string) bool
}

// 通过传入实现类型返回默认的接口实现，impl => 1. grpc(tcp、udp...)、 2. tcp
func NewAUTH(impl string) AUTH {
	switch impl {
	case constant.GPRC:
		return &AuthenticationGrpc{}
	case constant.TCP:
		return &AuthenticationTcp{}
	default:
		return &AuthenticationGrpc{}
	}
}

// 2、通过注入的方式直接返回实现的接口
func NewAuth(a AUTH) AUTH {
	return a
}
