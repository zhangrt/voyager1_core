package star

import (
	"context"

	"github.com/zhangrt/voyager1_core/auth/grpc/pb"
	"github.com/zhangrt/voyager1_core/auth/luna"
	"github.com/zhangrt/voyager1_core/util"
)

// 用户信息接口GRPC实现
type ClaimantGrpc struct{}

// 获取用户信息
func (claimant *ClaimantGrpc) GetUser(token string) (*luna.CustomClaims, error) {
	var err error
	claims := new(luna.CustomClaims)
	conn, client := GetGrpcClient()
	defer CloseConn(conn)
	result, err := client.GetUser(context.Background(), &pb.Authentication{
		Token: token,
	})
	if err != nil {
		return claims, err
	}
	claims = util.GrpcProtoUser2Claims(result)
	return claims, err
}

func (claimant *ClaimantGrpc) GetUserID(token string) string {
	var ID string
	claims, err := claimant.GetUser(token)
	if err != nil {
		ID = claims.ID
	}
	return ID
}

func (claimant *ClaimantGrpc) GetUserAuthorityId(token string) []string {
	var RoleIds []string
	claims, err := claimant.GetUser(token)
	if err != nil {
		RoleIds = claims.RoleIds
	}
	return RoleIds
}
