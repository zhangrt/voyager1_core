package star

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/zhangrt/voyager1_core/constant"
	pb "github.com/zhangrt/voyager1_core/zinx/pb"
	"google.golang.org/protobuf/proto"
)

type Message struct {
	Len   uint32
	MsgID uint32
	Data  []byte
}

// Client
type TcpClient struct {
	conn     net.Conn
	PID      int32
	isOnline bool
}

func (client *TcpClient) Unpack(headdata []byte) (head *Message, err error) {
	headBuf := bytes.NewReader(headdata)

	head = &Message{}

	// 读取Len
	if err = binary.Read(headBuf, binary.LittleEndian, &head.Len); err != nil {
		return nil, err
	}

	// 读取MsgID
	if err = binary.Read(headBuf, binary.LittleEndian, &head.MsgID); err != nil {
		return nil, err
	}

	// 封包太大
	//if head.Len > MaxPacketSize {
	//	return nil, packageTooBig
	//}

	return head, nil
}

func (client *TcpClient) Pack(msgID uint32, dataBytes []byte) (out []byte, err error) {
	outbuff := bytes.NewBuffer([]byte{})
	// 写Len
	if err = binary.Write(outbuff, binary.LittleEndian, uint32(len(dataBytes))); err != nil {
		return
	}
	// 写MsgID
	if err = binary.Write(outbuff, binary.LittleEndian, msgID); err != nil {
		return
	}

	//all pkg data
	if err = binary.Write(outbuff, binary.LittleEndian, dataBytes); err != nil {
		return
	}

	out = outbuff.Bytes()

	return
}

func (client *TcpClient) SendMsg(msgID uint32, data proto.Message) {

	// 进行编码
	binaryData, err := proto.Marshal(data)
	if err != nil {
		lg := fmt.Sprintf("marshaling error:  %s", err)
		fmt.Println(lg)
		return
	}

	sendData, err := client.Pack(msgID, binaryData)
	if err == nil {
		_, _ = client.conn.Write(sendData)
	} else {
		fmt.Println(err)
	}
}

/*
	处理一个回执业务
	获取Server回执的信息并进行业务处理
	根据不同的MsgID将结果存入不同的结果集中
*/
func (client *TcpClient) DoMsg(msg *Message) {
	//处理消息
	lg := fmt.Sprintf("msg ID :%d, data len: %d", msg.MsgID, msg.Len)
	fmt.Println(lg)
	switch msg.MsgID {
	case constant.TOKEN_RES: // token验证回执

		result := &pb.Result{}
		_ = proto.Unmarshal(msg.Data, result)

		println("TOKEN_RES : ", result.Success)
		println("TOKEN_RES :", result.Msg)

		// Key作为客户端验证结果与请求一一对应的唯一性标识(可用UUID)，需要在请求时传入，并不做任何处理，由Server返回
		StatelliteMgrObj.SetTokenResult(result.Key, result)
		// 必须做的步骤
		StatelliteMgrObj.RemoveMsgKey(result.Key)

	case constant.POLICY_RES: // 权限验证回执

		result := &pb.Result{}
		_ = proto.Unmarshal(msg.Data, result)

		StatelliteMgrObj.SetPolicyResult(result.Key, result)

	case constant.USER_RES: // 用户信息回执

		user := &pb.User{}
		_ = proto.Unmarshal(msg.Data, user)

		StatelliteMgrObj.SetUserResult(user.Key, user)

	case constant.HEARTBEAT_RES: // 心跳回执
		receipe := &pb.Receipe{}
		_ = proto.Unmarshal(msg.Data, receipe)
		client.PID = receipe.Id
		println("HEARTBEAT_RES ID:", receipe.Id)

		client.isOnline = true
	default:
		fmt.Println("unknown msg has received")
	}

}

func (client *TcpClient) Start() {

	// 处理回执
	go func() {
		for {

			//读取服务端发来的数据
			//1.读取8字节
			//第一次读取，读取数据头
			headData := make([]byte, 8)

			if _, err := io.ReadFull(client.conn, headData); err != nil {
				fmt.Println(err)
				return
			}
			pkgHead, err := client.Unpack(headData)
			if err != nil {
				return
			}
			//data
			if pkgHead.Len > 0 {
				pkgHead.Data = make([]byte, pkgHead.Len)
				if _, err := io.ReadFull(client.conn, pkgHead.Data); err != nil {
					return
				}
			}

			//处理服务器回执业务
			client.DoMsg(pkgHead)
		}
	}()

	// 心跳
	go func() {
		for {
			if client.conn == nil {
				break
			}
			client.HeartBeat()
			time.Sleep(time.Second * 5)
		}
	}()

	// 业务
	go func() {
		client.DoBusiness()
	}()

	// go func() {
	// 	for {
	// 		if client.isOnline {
	// 			fmt.Println("Online:::::::::::::::::::::::::::", client.isOnline)
	// 		} else {
	// 			fmt.Println("DisOnline :::::::::::::::::::::::::::", client.isOnline)
	// 			time.Sleep(time.Second)
	// 		}
	// 		time.Sleep(time.Second * 10)
	// 		// select {
	// 		// case <-client.isOnline:
	// 		// 	fmt.Println("Online:::::::::::::::::::::::::::", <-client.isOnline)
	// 		// 	return
	// 		// }
	// 		// case <-time.After(time.Second * 30):
	// 		// 	println("Close Conn >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	// 		// 	_ = client.conn.Close()
	// 		// 	return
	// 		// }
	// 	}
	// }()

}

// remote请求的业务协程
func (client *TcpClient) DoBusiness() {

	fmt.Println("business start .....")

	for {
		key, msgID := StatelliteMgrObj.GetMsgKey()
		if key == "" {
			time.Sleep(time.Millisecond * 10)
			continue
		} else {
			StatelliteMgrObj.RemoveMsgKey(key)
		}

		var msg proto.Message

		switch msgID {

		case constant.TOKEN_REQ:

			println("TOKEN_REQ :", key)
			msg = StatelliteMgrObj.GetTokenReq(key)

		case constant.POLICY_REQ:

			msg = StatelliteMgrObj.GetPolicyReq(key)

		case constant.USER_REQ:

			msg = StatelliteMgrObj.GetPolicyReq(key)

		default:
			println("unknown msgKey.msgID")
		}

		if msg != nil {
			client.SendMsg(msgID, msg)

		}

	}

}

// 心跳
func (client *TcpClient) HeartBeat() {
	msg := &pb.HearBeat{}
	msg.Id = client.PID
	println("heartBeat REQ ID:", msg.Id)
	client.SendMsg(constant.HEARTBEAT_REQ, msg)
}

func NewTcpClient(ip string, port int) *TcpClient {
	addrStr := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("tcp", addrStr)
	if err != nil {
		panic(err)
	}

	client := &TcpClient{
		conn:     conn,
		PID:      0,
		isOnline: false,
	}
	return client
}
