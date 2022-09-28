package star

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/zhangrt/voyager1_core/auth/grpc/pb"
	"github.com/zhangrt/voyager1_core/auth/luna"
	"github.com/zhangrt/voyager1_core/util"
	"google.golang.org/grpc"
)

// 用户信息接口GRPC实现
type ClaimantGrpc struct{}

// 获取用户信息
func (claimant *ClaimantGrpc) GetUser(token string) (*luna.CustomClaims, error) {
	var err error
	claims := new(luna.CustomClaims)
	conn, client := GetGrpcClient(grpc.WithInsecure())
	defer CloseConn(conn)
	result, err := client.GetUser(context.Background(), &pb.Token{
		Token: token,
	})
	if err != nil {
		return nil, err
	}
	claims = util.GrpcProtoUser2Claims(result)
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