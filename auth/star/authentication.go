package star

import "github.com/zhangrt/voyager1_core/auth/luna"

type Authentication struct{}

func (authentication *Authentication) ReadAuthentication(token string) (bool, string, *luna.CustomClaims) {
	var r bool
	var msg string
	var claims *luna.CustomClaims

	return r, msg, claims
}

func (authentication *Authentication) GrantedAuthority(path string, method string) bool {
	var r bool

	return r
}
