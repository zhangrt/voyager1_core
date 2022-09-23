package luna

import (
	"fmt"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

type Router struct {
	ID     uint32
	ROUTER ziface.IRouter
}

func Server(routers ...Router) ziface.IServer {
	s := znet.NewServer()
	s.SetOnConnStart(OnConnecionAdd)
	s.SetOnConnStop(OnConnectionLost)
	for _, r := range routers {
		s.AddRouter(r.ID, r.ROUTER)
	}
	return s
}

//当客户端建立连接的时候的hook函数
func OnConnecionAdd(conn ziface.IConnection) {
	//创建客户端Star
	star := NewStar(conn)

	//将当前新上线 star 添加到 starManager 中
	StarMgrObj.AddStar(star)

	//将该连接绑定属性PID
	conn.SetProperty("pID", star.PID)

	fmt.Println("=====> Star pIDID = ", star.PID, " arrived ====")
}

//当客户端断开连接的时候的hook函数
func OnConnectionLost(conn ziface.IConnection) {
	//获取当前连接的PID属性
	pID, _ := conn.GetProperty("pID")
	var starId int32
	if pID != nil {
		starId = pID.(int32)
	}

	//根据pID获取对应的 star 对象
	star := StarMgrObj.GetStarByPID(starId)

	//触发 star 下线业务
	if star != nil {
		star.LostConnection()
	}

	fmt.Println("====> Star ", starId, " left =====")

}
