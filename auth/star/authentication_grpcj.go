package star

import (
	"context"

	"github.com/zhangrt/voyager1_core/auth/luna"
	pb "github.com/zhangrt/voyager1_core/com/gs/voyager1_core/auth/grpc/pb"
	"github.com/zhangrt/voyager1_core/util"
)

// 授权鉴权接口Grpc的实现
type AuthenticationGrpcJ struct{}

func (authentication *AuthenticationGrpcJ) GrantedAuthority(token string, path string, method string) (bool, string, *luna.CustomClaims, string) {
	var r bool
	conn, client := GetGrpcClientJ()
	defer CloseConn(conn)
	result, err := client.GrantedAuthority(context.Background(), &pb.Policy{
		Token:  token,
		Path:   path,
		Method: method,
	})
	if err != nil {
		r = false
		return r, err.Error(), nil, ""
	} else {
		return result.Success, result.Msg, util.GrpcProtoClaimsTransformClaimsJ(result.Claims), result.NewToken
	}
}
