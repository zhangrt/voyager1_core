package star

import (
	"context"

	"github.com/zhangrt/voyager1_core/auth/grpc/pb"
	"github.com/zhangrt/voyager1_core/auth/luna"
	util "github.com/zhangrt/voyager1_core/util"
)

// 授权鉴权接口实现
type AuthenticationGrpc struct{}

func (authentication *AuthenticationGrpc) ReadAuthentication(token string) (bool, string, *luna.CustomClaims) {
	var r bool
	var msg string
	var claims *luna.CustomClaims
	conn, client := GetGrpcClient()
	defer CloseConn(conn)
	result, err := client.ReadAuthentication(context.Background(), &pb.Token{
		Token: token,
	})
	if err != nil {
		msg = err.Error()
	} else {
		r = result.Success
		msg = result.Msg
		claims = util.GrpcProtoResult2Claims(result)
	}
	return r, msg, claims
}

func (authentication *AuthenticationGrpc) GrantedAuthority(authorityId string, path string, method string) bool {
	var r bool
	conn, client := GetGrpcClient()
	defer CloseConn(conn)
	result, err := client.GrantedAuthority(context.Background(), &pb.Policy{
		AuthorityId: authorityId,
		Path:        path,
		Method:      method,
	})
	if err != nil {
		r = false
	} else {
		r = result.Success
	}
	return r
}
