package core

import (
	"fmt"
	"sync"

	"github.com/aceld/zinx/ziface"
	"google.golang.org/protobuf/proto"
)

type Star struct {
	PID  int32              // 客户端 star ID
	Conn ziface.IConnection // 当前 star 的连接
}

/*
	Star ID 生成器
*/
var PIDGen int32 = 1  // 用来生成 star ID的计数器
var IDLock sync.Mutex // 保护PIDGen的互斥机制

func NewStar(conn ziface.IConnection) *Star {
	//生成一个PID
	IDLock.Lock()
	ID := PIDGen
	PIDGen++
	IDLock.Unlock()

	s := &Star{
		PID:  ID,
		Conn: conn,
	}

	return s
}

// star 下线
func (s *Star) LostConnection() {

	StarMgrObj.RemoveStarByPID(s.PID)
}

/*
	发送消息给客户端，
	主要是将pb的protobuf数据序列化之后发送
*/
func (s *Star) SendMsg(msgID uint32, data proto.Message) {
	//fmt.Printf("before Marshal data = %+v\n", data)
	// 将proto Message结构体序列化
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("marshal msg err: ", err)
		return
	}
	//f mt.Printf("after Marshal data = %+v\n", msg)

	if s.Conn == nil {
		fmt.Println("connection in Star is nil")
		return
	}

	// 调用Zinx框架的SendMsg发包
	if err := s.Conn.SendMsg(msgID, msg); err != nil {
		fmt.Println("Star SendMsg error !")
		return
	}

	return
}

func (s *Star) Check(token string) {

}

func (s *Star) AuthenticationRequest(path string, method string) {

}
