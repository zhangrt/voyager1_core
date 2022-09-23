package api

import (
	"fmt"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/golang/protobuf/proto"
	"github.com/zhangrt/voyager1_core/zinx/core/luna"
	pb "github.com/zhangrt/voyager1_core/zinx/pb"
)

// 鉴权
type AuthenticationRequestApi struct {
	znet.BaseRouter
}

func (*AuthenticationRequestApi) Handle(request ziface.IRequest) {
	//1. 将客户端传来的proto协议解码
	msg := &pb.Policy{}
	err := proto.Unmarshal(request.GetData(), msg)
	if err != nil {
		fmt.Println("Move: Position Unmarshal error ", err)
		return
	}

	//2. 得知当前的消息是从哪个 star 传递来的,从连接属性pID中获取
	pID, err := request.GetConnection().GetProperty("pID")
	if err != nil {
		fmt.Println("GetProperty pID error", err)
		request.GetConnection().Stop()
		return
	}

	//3. 根据pID得到 star 对象
	star := luna.StarMgrObj.GetStarByPID(pID.(int32))

	//4. 让 star 对象鉴权
	star.AuthenticationRequest(msg)
}
