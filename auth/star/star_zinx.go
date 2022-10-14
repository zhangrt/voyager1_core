package star

import (
	uuid "github.com/satori/go.uuid"
	"github.com/zhangrt/voyager1_core/constant"
	"github.com/zhangrt/voyager1_core/global"
	"github.com/zhangrt/voyager1_core/zinx/core/star"
	pb "github.com/zhangrt/voyager1_core/zinx/pb"
)

// 需要重构 增加错误重试重连等机制
func StartClient() {
	client := star.NewTcpClient(global.G_CONFIG.Zinx.Host, global.G_CONFIG.Zinx.TcpPort)

	star.StatelliteMgrObj.ClientObj["auth"] = client
	go client.Start()

}

func SendProtoTokenMsg(token string, msgId uint32) string {
	key := uuid.NewV4().String()
	star.StatelliteMgrObj.AddTokenReq(key, &pb.Token{
		Key:   key,
		Token: token,
	})
	star.StatelliteMgrObj.AddMsgKey(key, msgId)
	return key
}

func SendProtoPolicyMsg(token string, p string, m string, msgId uint32) string {
	key := uuid.NewV4().String()
	star.StatelliteMgrObj.AddPolicyReq(key, &pb.Policy{
		Key:    key,
		Token:  token,
		Path:   p,
		Method: m,
	})
	star.StatelliteMgrObj.AddMsgKey(key, msgId)
	return key
}

// 清除Proto Msg 信息
func CleanProtoMsg(key string, msgId uint32) {
	switch msgId {
	case constant.TOKEN_REQ:
		star.StatelliteMgrObj.RemoveMsgKey(key)
		star.StatelliteMgrObj.RemoveTokenResult(key)
	case constant.USER_REQ:
		star.StatelliteMgrObj.RemoveMsgKey(key)
		star.StatelliteMgrObj.RemoveUserResult(key)
	case constant.POLICY_REQ:
		star.StatelliteMgrObj.RemoveMsgKey(key)
		star.StatelliteMgrObj.RemovePolicyResult(key)
	}
}

func RemoteTimeout(key string) string {
	star.StatelliteMgrObj.RemoveMsgKey(key)
	return "Timeout remote request for ReadAuthentication"
}
