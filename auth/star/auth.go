package star

import "github.com/xyy277/gallery/auth/luna"

// star 请求 luna的接口进行鉴权与授权
type AUTH interface {
	// 通过token请求鉴权
	RemoteAuthentication(token string) (bool, string, *luna.CustomClaims)
	// 通过request中的path和method请求验证策略
	RemoteAuthenticationPolicy(path string, method string) bool
}

func NewAUTH() AUTH {
	// Request impl
	return &Authentication{}
}
