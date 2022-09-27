package star

import "github.com/zhangrt/voyager1_core/auth/luna"

// 授权鉴权接口实现
type AuthenticationGrpc struct{}

func (authentication *AuthenticationGrpc) ReadAuthentication(token string) (bool, string, *luna.CustomClaims) {
	var r bool
	var msg string
	var claims *luna.CustomClaims
	return r, msg, claims
}

func (authentication *AuthenticationGrpc) GrantedAuthority(authorityId string, path string, method string) bool {
	var r bool

	return r
}
