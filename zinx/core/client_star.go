package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"

	pb "github.com/zhangrt/voyager1_core/zinx/pb"
	"google.golang.org/protobuf/proto"
)

type Message struct {
	Len   uint32
	MsgID uint32
	Data  []byte
}

type TcpClient struct {
	conn     net.Conn
	PID      int32
	isOnline chan bool

	token   chan pb.Token
	police  chan pb.Police
	hasAuth chan bool
	claims  chan pb.Result
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
		fmt.Println(fmt.Sprintf("marshaling error:  %s", err))
		return
	}

	sendData, err := client.Pack(msgID, binaryData)
	if err == nil {
		_, _ = client.conn.Write(sendData)
	} else {
		fmt.Println(err)
	}

	return
}

/*
	处理一个回执业务
*/
func (client *TcpClient) DoMsg(msg *Message) {
	//处理消息
	fmt.Println(fmt.Sprintf("msg ID :%d, data len: %d", msg.MsgID, msg.Len))
	if msg.MsgID == 1 {
		//服务器回执给客户端 分配ID

		//解析proto
		syncpID := &pb.Result{}
		_ = proto.Unmarshal(msg.Data, syncpID)
		client.isOnline <- true
		//给当前客户端ID进行赋值
	} else if msg.MsgID == 2 {
		//服务器回执客户端广播数据

		//解析proto
		bdata := &pb.User{}
		_ = proto.Unmarshal(msg.Data, bdata)
		client.isOnline <- true

	}
}

func (client *TcpClient) Start() {
	go func() {
		for {
			//读取服务端发来的数据 ==》 SyncPID
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

	// 10s后，断开连接
	for {
		select {
		case <-client.isOnline:
			go func() {
				for {
					// TODO something
					time.Sleep(time.Second)
				}
			}()
		case <-time.After(time.Second * 10):
			_ = client.conn.Close()
			return
		}
	}
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
		isOnline: make(chan bool),

		hasAuth: make(chan bool),

		token:  make(chan pb.Token),
		police: make(chan pb.Police),
		claims: make(chan pb.Result),
	}
	return client
}
