package luna

import (
	"time"

	redis "github.com/xyy277/gallery/cache"
	"github.com/xyy277/gallery/global"
)

type JwtService struct{}

//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func (jwtService *JwtService) JsonInBlacklist(jwtList JwtBlacklist) (err error) {
	err = global.G_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
	// err := global.GS_DB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	// isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	// return !isNotFound
}

//@function: GetCacheJWT
//@description: 从redis取jwt
//@param: userName string
//@return: redisJWT string, err error

func (jwtService *JwtService) GetCacheJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = redis.GetResult(userName)
	return redisJWT, err
}

//@function: SetCacheJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error

func (jwtService *JwtService) SetCacheJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.G_CONFIG.JWT.ExpiresTime) * time.Second
	err = redis.Set(userName, jwt, timer)
	return err
}
