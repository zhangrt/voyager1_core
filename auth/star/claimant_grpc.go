package star

import (
	uuid "github.com/satori/go.uuid"
	"github.com/zhangrt/voyager1_core/auth/luna"
)

// 用户信息接口GRPC实现
type ClaimantGrpc struct{}

// 获取用户信息
func (claimant *ClaimantGrpc) GetUser(token string) (*luna.CustomClaims, error) {
	var err error
	claims := new(luna.CustomClaims)

	return claims, err
}

func (claimant *ClaimantGrpc) GetUserID(token string) uint {
	var ID uint
	claims, err := claimant.GetUser(token)
	if err != nil {
		ID = claims.ID
	}
	return ID
}

func (claimant *ClaimantGrpc) GetUserUUID(token string) uuid.UUID {
	var UUID uuid.UUID
	claims, err := claimant.GetUser(token)
	if err != nil {
		UUID = claims.UUID
	}
	return UUID
}

func (claimant *ClaimantGrpc) GetUserAuthorityId(token string) string {
	var AuthorityId string
	claims, err := claimant.GetUser(token)
	if err != nil {
		AuthorityId = claims.AuthorityId
	}
	return AuthorityId
}
