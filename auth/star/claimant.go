package star

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xyy277/gallery/auth/luna"
)

type Claimant struct{}

func (claimant *Claimant) GetClaims(token string) (*luna.CustomClaims, error) {
	var claims *luna.CustomClaims
	var err error

	return claims, err
}
func (claimant *Claimant) GetUserID(token string) uint {
	var ID uint

	return ID
}
func (claimant *Claimant) GetUserUUID(token string) uuid.UUID {
	var UUID uuid.UUID

	return UUID
}
func (claimant *Claimant) GetUserAuthorityId(token string) string {
	var AuthorityId string

	return AuthorityId
}
