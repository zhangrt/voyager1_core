package star

import (
	"bytes"
	"container/list"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"runtime"
	"sync"
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
	conn        net.Conn
	PID         int32
	isOnline    chan bool
	RequestKeys *list.List
}

var keyLock sync.RWMutex

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

		println("TokenReq Success: ", result.Success)
		println("TokenReq Msg:", result.Msg)

		// Key作为客户端验证结果与请求一一对应的唯一性标识(可用UUID)，需要在请求时传入，并不做任何处理，由Server返回
		StatelliteMgrObj.SetTokenResult(result.Key, result)

		// 必须做的步骤
		client.RemoveKey(MsgKey{
			Key:   result.Key,
			MsgID: msg.MsgID,
		})
		// 解锁
		keyLock.RUnlock()

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
		client.isOnline <- true
	default:
		fmt.Println("unknown msg has received")
	}

}

func (client *TcpClient) Start() {
	// 处理回执
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

	// 心跳
	go func() {
		for {
			client.HeartBeat()
			time.Sleep(time.Second * 5)
		}
	}()

	// 业务
	go func() {
		client.DoBusiness()
	}()

	// 10s后，断开连接
	for {
		select {
		case <-client.isOnline:
			fmt.Println("heartbeat")
		case <-time.After(time.Second * 10):
			_ = client.conn.Close()
			return
		}
	}
}

// remote请求的业务协程
func (client *TcpClient) DoBusiness() {
	fmt.Println("business start .....")
	var wg sync.WaitGroup
	num := runtime.NumCPU()
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {

			for {
				// 从 msg key 数据集中获取存在的数据信息
				if client.hasKey() {
					// 锁
					keyLock.RLock()
					msgKey := client.getKey()
					var msg proto.Message

					switch msgKey.MsgID {

					case constant.TOKEN_REQ:

						msg = StatelliteMgrObj.GetTokenReq(msgKey.Key)

					case constant.POLICY_REQ:

						msg = StatelliteMgrObj.GetPolicyReq(msgKey.Key)

					case constant.USER_REQ:

						msg = StatelliteMgrObj.GetPolicyReq(msgKey.Key)

					default:
						println("unknown msgKey.msgID")
					}

					if msg != nil {
						client.SendMsg(msgKey.MsgID, msg)
					}

				}

			}
			wg.Done()

		}()
	}
	wg.Wait()

}

// 心跳
func (client *TcpClient) HeartBeat() {
	msg := &pb.HearBeat{}
	msg.Id = client.PID
	println("heartBeat ID:", msg.Id)
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
		isOnline: make(chan bool),
	}
	client.RequestKeys = list.New()
	return client
}

// 采用pushback
func (client *TcpClient) PushKey(msgKey MsgKey) {
	keyLock.Lock()
	l := client.RequestKeys
	l.PushBack(msgKey)
	keyLock.Unlock()
}

func (client *TcpClient) RemoveKey(msgKey MsgKey) {
	keyLock.Lock()
	l := client.RequestKeys
	for e := l.Front(); e != nil; {
		next := e.Next()
		// 查找删除
		if e.Value.(MsgKey).Key == msgKey.Key {
			l.Remove(e)
			break
		}
		e = next
	}
	keyLock.Unlock()
}

// get front
func (client *TcpClient) getKey() MsgKey {
	keyLock.RLock()
	e := client.RequestKeys.Front()
	defer keyLock.RUnlock()
	if e == nil {
		return MsgKey{}
	}
	return e.Value.(MsgKey)
}

func (client *TcpClient) hasKey() bool {
	return client.RequestKeys.Len() > 0
}
