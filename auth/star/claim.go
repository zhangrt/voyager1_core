package star

import (
	uuid "github.com/satori/go.uuid"
	"github.com/zhangrt/voyager1_core/auth/luna"
)

type CLAIM interface {
	GetClaims(token string) (*luna.CustomClaims, error)
	GetUserID(token string) uint
	GetUserUUID(token string) uuid.UUID
	GetUserAuthorityId(token string) string
}

func NewCLAMI() CLAIM {

	return &Claimant{}
}
