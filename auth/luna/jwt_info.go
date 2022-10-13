package luna

import (
	"github.com/zhangrt/voyager1_core/global"

	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

type JwtBlacklist struct {
	global.GS_BASE_MODEL_ID_UINT
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
	// UUID 标准通用的uuid
	UUID uuid.UUID
	// ID 一般对应于数组库默认的自增主键
	ID uint
	// 用户名
	Account string
	Phone   string
	Email   string
	// 名称
	Name string
	// // 权限ID
	// RoleId  string
	RoleIds []string
	// // 权限信息
	// Role interface{}
	// 权限信息
	Roles interface{}
	// 部门Id
	DepartMentId string
	// 部门名称
	DepartMentName string
	// 组织机构ID
	OrganizationId string
	// 组织机构名称
	OrganizationName string
}
