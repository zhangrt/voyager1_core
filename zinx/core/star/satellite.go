package star

import (
	"sync"

	pb "github.com/zhangrt/voyager1_core/zinx/pb"
)

/*
 Statellite 管理remote 请求和返回的数据的结果集
 并发安全
 pb 中的key为唯一标识
*/
type Statellite struct {
	TokenReq      map[string]*pb.Authentication // token验证
	PolicyReq     map[string]*pb.Policy         // 权限校验
	UserReq       map[string]*pb.Authentication // 用户信息获取
	ResultToken   map[string]*pb.Result         // token验证结果
	ResultPolicy  map[string]*pb.Result         // 权限校验结果
	ResultUser    map[string]*pb.User           // 用户信息获取结果
	MsgKeys       map[string]uint32             // key - MsgID 存放唯一键值对key和请求ID
	keyLock       sync.RWMutex
	setTokenLock  sync.RWMutex
	setPolicyLock sync.RWMutex
	setUserLock   sync.RWMutex
	getTokenLog   sync.RWMutex
	getPolicyLock sync.RWMutex
	getUserLock   sync.RWMutex
	ClientObj     map[string]*TcpClient
}

var (
	// 提供一个对外的 Statellite 管理句柄，管理所有发送和接收数据
	StatelliteMgrObj *Statellite
)

//提供 Statellite 初始化方法
func init() {
	StatelliteMgrObj = &Statellite{
		MsgKeys:      map[string]uint32{},
		TokenReq:     make(map[string]*pb.Authentication),
		PolicyReq:    make(map[string]*pb.Policy),
		UserReq:      make(map[string]*pb.Authentication),
		ResultToken:  make(map[string]*pb.Result), // token验证结果
		ResultPolicy: make(map[string]*pb.Result), // 权限校验结果
		ResultUser:   make(map[string]*pb.User),   // 用户信息结果
		ClientObj:    make(map[string]*TcpClient),
	}
}

func (statellite *Statellite) AddMsgKey(key string, msgID uint32) {
	statellite.keyLock.Lock()
	statellite.MsgKeys[key] = msgID
	statellite.keyLock.Unlock()
}

func (statellite *Statellite) RemoveMsgKey(key string) {
	statellite.keyLock.Lock()
	delete(statellite.MsgKeys, key)
	statellite.keyLock.Unlock()
}

func (statellite *Statellite) GetMsgKey() (string, uint32) {
	statellite.keyLock.RLock()
	k := ""
	var m uint32
	ms := statellite.MsgKeys
	if len(ms) > 0 {
		for e := range ms {
			k = e
			m = ms[k]
		}
	}
	statellite.keyLock.RUnlock()
	return k, m
}

func (statellite *Statellite) GetMsgIDByKey(key string) uint32 {
	statellite.keyLock.Lock()
	id := statellite.MsgKeys[key]
	defer statellite.keyLock.Unlock()
	return id
}

// --------------------------------------------------------------------------Result----------------------------------------------------------------------------
// --------------------------------token校验结果-----------------------------------
func (statellite *Statellite) SetTokenResult(key string, pb *pb.Result) {
	statellite.getTokenLog.Lock()
	pb.Key = key
	statellite.ResultToken[key] = pb
	statellite.getTokenLog.Unlock()
}

func (statellite *Statellite) RemoveTokenResult(key string) {
	statellite.getTokenLog.Lock()
	delete(statellite.ResultToken, key)
	statellite.getTokenLog.Unlock()
}

func (statellite *Statellite) GetTokenResult(key string) *pb.Result {
	statellite.getTokenLog.RLock()
	defer statellite.getTokenLog.RUnlock()

	return statellite.ResultToken[key]
}

// --------------------------------policy校验结果-----------------------------------
func (statellite *Statellite) SetPolicyResult(key string, pb *pb.Result) {
	statellite.getPolicyLock.Lock()
	pb.Key = key
	statellite.ResultToken[key] = pb
	statellite.getPolicyLock.Unlock()
}

func (statellite *Statellite) RemovePolicyResult(key string) {
	statellite.getPolicyLock.Lock()
	delete(statellite.ResultToken, key)
	statellite.getPolicyLock.Unlock()
}

func (statellite *Statellite) GetPolicyResult(key string) *pb.Result {
	statellite.getPolicyLock.RLock()
	defer statellite.getPolicyLock.RUnlock()

	return statellite.ResultPolicy[key]
}

// --------------------------------user校验结果-----------------------------------
func (statellite *Statellite) SetUserResult(key string, pb *pb.User) {
	statellite.getUserLock.Lock()
	pb.Key = key
	statellite.ResultUser[key] = pb
	statellite.getUserLock.Unlock()
}

func (statellite *Statellite) RemoveUserResult(key string) {
	statellite.getUserLock.Lock()
	delete(statellite.ResultToken, key)
	statellite.getUserLock.Unlock()
}

func (statellite *Statellite) GetUserResult(key string) *pb.User {
	statellite.getUserLock.RLock()
	defer statellite.getUserLock.RUnlock()

	return statellite.ResultUser[key]
}

// -------------------------------------------------------------------------Request---------------------------------------------------------------------------
// --------------------------------token-----------------------------------
func (statellite *Statellite) AddTokenReq(key string, pb *pb.Authentication) {
	statellite.setTokenLock.Lock()
	pb.Key = key
	statellite.TokenReq[key] = pb
	statellite.setTokenLock.Unlock()
}

func (statellite *Statellite) RemoveStarByPID(key string) {
	statellite.setTokenLock.Lock()
	delete(statellite.TokenReq, key)
	statellite.setTokenLock.Unlock()
}

func (statellite *Statellite) GetTokenReq(key string) *pb.Authentication {
	statellite.setTokenLock.RLock()
	defer statellite.setTokenLock.RUnlock()

	return statellite.TokenReq[key]
}

// --------------------------------policy----------------------------------
func (statellite *Statellite) AddPolicyReq(key string, pb *pb.Policy) {
	statellite.setPolicyLock.Lock()
	pb.Key = key
	statellite.PolicyReq[key] = pb
	statellite.setPolicyLock.Unlock()
}

func (statellite *Statellite) RemovePolicyReq(key string) {
	statellite.setPolicyLock.Lock()
	delete(statellite.PolicyReq, key)
	statellite.setPolicyLock.Unlock()
}

func (statellite *Statellite) GetPolicyReq(key string) *pb.Policy {
	statellite.setPolicyLock.RLock()
	defer statellite.setPolicyLock.RUnlock()

	return statellite.PolicyReq[key]
}

// --------------------------------user-----------------------------------
func (statellite *Statellite) AddUserReq(key string, pb *pb.Authentication) {
	statellite.setUserLock.Lock()
	pb.Key = key
	statellite.UserReq[key] = pb
	statellite.setUserLock.Unlock()
}

func (statellite *Statellite) RemoveUserReq(key string) {
	statellite.setUserLock.Lock()
	delete(statellite.UserReq, key)
	statellite.setUserLock.Unlock()
}

func (statellite *Statellite) GetUserReq(key string) *pb.Authentication {
	statellite.setUserLock.RLock()
	defer statellite.setUserLock.RUnlock()

	return statellite.UserReq[key]
}
