package star

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xyy277/gallery/auth/luna"
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
