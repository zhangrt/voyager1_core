package star

import (
	"time"

	"github.com/zhangrt/voyager1_core/auth/luna"
	util "github.com/zhangrt/voyager1_core/util"

	"github.com/zhangrt/voyager1_core/constant"
	"github.com/zhangrt/voyager1_core/zinx/core/star"
	pb "github.com/zhangrt/voyager1_core/zinx/pb"
)

// 授权鉴权接口实现
type AuthenticationTcp struct{}

func (authentication *AuthenticationTcp) ReadAuthentication(token string) (bool, string, *luna.CustomClaims) {
	var msg string
	var claims *luna.CustomClaims
	key := SendProtoTokenMsg(token, constant.TOKEN_REQ)
	// 设置超时时间
	timeout := time.After(time.Second * 10)
	var result *pb.Result
	for {
		// 结果
		result = star.StatelliteMgrObj.GetTokenResult(key)
		if result != nil {
			println("Key >>>>> ", key)
			println("<GetTokenResult Key>: ", result.Key)
			println("====================== %d =========================", key == result.Key)
			CleanProtoMsg(key, constant.TOKEN_REQ)
			break
		}
		time.Sleep(time.Millisecond * 5)
		select {
		case <-timeout:
			msg = RemoteTimeout(key)
			return false, msg, nil
		}

	}
	msg = "ReadAuthentication Success"

	claims = util.ZinxProtoResult2Claims(result)

	return true, msg, claims
}

func (authentication *AuthenticationTcp) GrantedAuthority(authorityId string, path string, method string) bool {
	var r bool
	key := SendProtoPolicyMsg(authorityId, path, method, constant.POLICY_REQ)

	// 设置超时时间
	timeout := time.After(time.Second * 10)
	var result *pb.Result
	for {
		// 结果
		result = star.StatelliteMgrObj.GetPolicyResult(key)
		if result != nil {
			println("Key >>>>> ", key)
			println("<GetPolickResult Key>: ", result.Key)
			println("====================== %d =========================", key == result.Key)
			CleanProtoMsg(key, constant.POLICY_REQ)

			r = true
			break
		}
		time.Sleep(time.Millisecond * 5)

		select {
		case <-timeout:
			RemoteTimeout(key)
			r = false
			return r
		}

	}
	return r
}
