package star

import (
	uuid "github.com/satori/go.uuid"
	"github.com/zhangrt/voyager1_core/auth/luna"
	"github.com/zhangrt/voyager1_core/constant"
)

type CLAIM interface {
	GetUser(token string) (*luna.CustomClaims, error)
	GetUserID(token string) uint
	GetUserUUID(token string) uuid.UUID
	GetUserAuthorityId(token string) string
}

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
