package luna

import (
	"github.com/xyy277/gallery/global"

	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

type JwtBlacklist struct {
	global.GS_BASE_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

// v0.2 针对基础业务做补充，增加部门id、单位id等
type BaseClaims struct {
	// UUID
	UUID uuid.UUID
	// ID
	ID uint
	// 用户名
	Username string
	// 昵称
	NickName string
	// 权限ID
	AuthorityId string
	// 权限信息
	Authority interface{}
	// 部门Id
	DepartMentId string
	// 部门名称
	DepartMentName string
	// 单位Id
	UnitId string
	// 单位名称
	UnitName string
}
