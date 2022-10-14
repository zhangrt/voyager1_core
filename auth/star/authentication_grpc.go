package star

import (
	"context"

	"github.com/zhangrt/voyager1_core/auth/grpc/pb"
	"github.com/zhangrt/voyager1_core/auth/luna"
	"github.com/zhangrt/voyager1_core/util"
)

// 授权鉴权接口Grpc的实现
type AuthenticationGrpc struct{}

func (authentication *AuthenticationGrpc) GrantedAuthority(token string, path string, method string) (bool, string, *luna.CustomClaims, string) {
	var r bool
	conn, client := GetGrpcClient()
	defer CloseConn(conn)
	result, err := client.GrantedAuthority(context.Background(), &pb.Policy{
		Token:  token,
		Path:   path,
		Method: method,
	})
	if err != nil {
		r = false
	} else {
		r = result.Success
	}
	return r, result.Msg, util.GrpcProtoClaimsTransformClaims(result.Claims), result.NewToken
}
