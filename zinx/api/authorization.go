package api

import (
	"fmt"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/golang/protobuf/proto"
	"github.com/zhangrt/voyager1_core/zinx/core"
	pb "github.com/zhangrt/voyager1_core/zinx/pb"
)

//授权 路由业务
type AuthorizationApi struct {
	znet.BaseRouter
}

func (*AuthorizationApi) Handle(request ziface.IRequest) {
	//1. 将客户端传来的proto协议解码
	msg := &pb.Token{}
	err := proto.Unmarshal(request.GetData(), msg)
	if err != nil {
		fmt.Println("Talk Unmarshal error ", err)
		return
	}

	//2. 得知当前的消息是从哪个玩家传递来的,从连接属性pID中获取
	pID, err := request.GetConnection().GetProperty("pID")
	if err != nil {
		fmt.Println("GetProperty pID error", err)
		request.GetConnection().Stop()
		return
	}
	//3. 根据pID得到 star 对象
	star := core.StarMgrObj.GetStarByPID(pID.(int32))

	//4. 让star对象验证token
	star.Check(msg.Token)
}
