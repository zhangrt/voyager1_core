package star

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/zhangrt/voyager1_core/auth/luna"
	"github.com/zhangrt/voyager1_core/constant"
	util "github.com/zhangrt/voyager1_core/util"
	"github.com/zhangrt/voyager1_core/zinx/core/star"
	pb "github.com/zhangrt/voyager1_core/zinx/pb"
)

// 用户信息接口TCP实现
type ClaimantZinx struct{}

func (claimant *ClaimantZinx) GetUser(token string) (*luna.CustomClaims, error) {
	var claims *luna.CustomClaims
	var err error
	key := SendProtoTokenMsg(token, constant.USER_REQ)
	// 设置超时时间
	timeout := time.After(time.Second * 10)
	var user *pb.User
	for {
		// 结果
		user = star.StatelliteMgrObj.GetUserResult(key)
		if user != nil {
			println("Key >>>>> ", key)
			println("<GetUserResult Key>: ", user.Key)
			println("====================== %d =========================", key == user.Key)
			CleanProtoMsg(key, constant.USER_REQ)

			break
		}

		time.Sleep(time.Millisecond * 5)

		select {
		case <-timeout:
			RemoteTimeout(key)
			return nil, err
		}
	}
	claims = util.ZinxProtoUser2Claims(user)

	return claims, err
}

func (claimant *ClaimantZinx) GetUserID(token string) string {
	var ID string
	claims, err := claimant.GetUser(token)
	if err != nil {
		ID = claims.ID.String()
	}
	return ID
}

func (claimant *ClaimantZinx) GetUserUUID(token string) uuid.UUID {
	var UUID uuid.UUID
	claims, err := claimant.GetUser(token)
	if err != nil {
		UUID = claims.ID
	}
	return UUID
}

func (claimant *ClaimantZinx) GetUserAuthorityId(token string) []string {
	var RoleIds []string
	claims, err := claimant.GetUser(token)
	if err != nil {
		RoleIds = claims.RoleIds
	}
	return RoleIds
}
