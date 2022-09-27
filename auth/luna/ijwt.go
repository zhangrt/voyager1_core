package luna

import (
	"fmt"
)

// 权限接口
type JWT interface {
	JsonInBlacklist(jwtList JwtBlacklist) (err error)
	IsBlacklist(jwt string) bool
	GetCacheJWT(userName string) (redisJWT string, err error)
	SetCacheJWT(jwt string, userName string) (err error)
	LoadAll() (err error)
}

var j JWT

// @author: [zhoujiajun](zhoujiajun@gsafety.com)
// 权限实现
func NewJWT() JWT {

	if j != nil {
		return j
	}
	fmt.Errorf("Interface JWT not implemented")
	return j
}

type UnimplementedJwt struct {
}

func (UnimplementedJwt) JsonInBlacklist(jwtList JwtBlacklist) (err error) {
	return fmt.Errorf("method JsonInBlacklist not implemented")
}

func (UnimplementedJwt) IsBlacklist(jwt string) bool {
	return false
}

func (UnimplementedJwt) GetCacheJWT(userName string) (redisJWT string, err error) {
	return "", fmt.Errorf("method GetCacheJWT not implemented")
}

func (UnimplementedJwt) SetCacheJWT(jwt string, userName string) (err error) {
	return fmt.Errorf("method SetCacheJWT not implemented")
}

func (UnimplementedJwt) LoadAll() error {
	return fmt.Errorf("method LoadAll not implemented")
}

func RegisterJwt(jw JWT) {
	j = jw
}
