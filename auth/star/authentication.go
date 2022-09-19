package star

import "github.com/xyy277/gallery/auth/luna"

type Authentication struct{}

func (authentication *Authentication) RemoteAuthentication(token string) (bool, string, *luna.CustomClaims) {
	var r bool
	var msg string
	var claims *luna.CustomClaims

	return r, msg, claims
}

func (authentication *Authentication) RemoteAuthenticationPolicy(path string, method string) bool {
	var r bool

	return r
}
