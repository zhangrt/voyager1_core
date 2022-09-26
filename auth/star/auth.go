package star

import "github.com/zhangrt/voyager1_core/auth/luna"

// star 请求 luna的接口进行鉴权与授权
type AUTH interface {
	// 通过token请求鉴权
	ReadAuthentication(token string) (bool, string, *luna.CustomClaims)
	// 通过request中的path和method请求验证策略
	GrantedAuthority(authorityId string, path string, method string) bool
}

func NewAUTH() AUTH {
	// Request impl
	return &Authentication{}
}
