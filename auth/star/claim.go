package star

import (
	uuid "github.com/satori/go.uuid"
	"github.com/zhangrt/voyager1_core/auth/luna"
	"github.com/zhangrt/voyager1_core/constant"
)

// 获取用户信息接口
type CLAIM interface {
	GetUser(token string) (*luna.CustomClaims, error)
	GetUserID(token string) uint
	GetUserUUID(token string) uuid.UUID
	GetUserAuthorityId(token string) string
}

// 通过传入实现类型返回不同的接口实现
func NewCLAMI(impl string) CLAIM {

	switch impl {
	case constant.GPRC:
		return &ClaimantGrpc{}
	case constant.TCP:
		return &ClaimantTcp{}
	default:
		return &ClaimantGrpc{}
	}

}

// 直接返回接口实现
func NewClami(c CLAIM) CLAIM {
	return c
}
