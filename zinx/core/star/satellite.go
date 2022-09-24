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
	TokenReq      map[string]*pb.Token  // token验证
	PolicyReq     map[string]*pb.Policy // 权限校验
	UserReq       map[string]*pb.Token  // 用户信息获取
	ResultToken   map[string]*pb.Result // token验证结果
	ResultPolicy  map[string]*pb.Result // 权限校验结果
	ResultUser    map[string]*pb.User   // 用户信息获取结果
	MsgKeys       map[string]uint32     // key - MsgID 存放唯一键值对key和请求ID
	setTokenLock  sync.RWMutex
	setPolicyLock sync.RWMutex
	setUserLock   sync.RWMutex
	getTokenLog   sync.RWMutex
	getPolicyLock sync.RWMutex
	getUserLock   sync.RWMutex
	ClientObj     map[int32]*TcpClient
}

var (
	StatelliteMgrObj *Statellite
)

//提供 Statellite 初始化方法
func init() {
	StatelliteMgrObj = &Statellite{
		MsgKeys:      map[string]uint32{},
		TokenReq:     make(map[string]*pb.Token),
		PolicyReq:    make(map[string]*pb.Policy),
		UserReq:      make(map[string]*pb.Token),
		ResultToken:  make(map[string]*pb.Result), // token验证结果
		ResultPolicy: make(map[string]*pb.Result), // 权限校验结果
		ResultUser:   make(map[string]*pb.User),   // 用户信息结果
		ClientObj:    make(map[int32]*TcpClient),
	}
}

func (statellite *Statellite) AddMsgKey(key string, msgID uint32) {

	statellite.MsgKeys[key] = msgID

}

func (statellite *Statellite) RemoveMsgKey(key string) {
	delete(statellite.MsgKeys, key)
}

func (statellite *Statellite) GetMsgKey() (string, uint32) {
	k := ""
	var m uint32
	ms := statellite.MsgKeys
	if len(ms) > 0 {
		for e := range ms {
			k = e
			m = ms[k]
		}
	}
	return k, m
}

func (statellite *Statellite) HasMsgKey() bool {

	return len(statellite.MsgKeys) > 0
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
func (statellite *Statellite) AddTokenReq(key string, pb *pb.Token) {
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

func (statellite *Statellite) GetTokenReq(key string) *pb.Token {
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
func (statellite *Statellite) AddUserReq(key string, pb *pb.Token) {
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

func (statellite *Statellite) GetUserReq(key string) *pb.Token {
	statellite.setUserLock.RLock()
	defer statellite.setUserLock.RUnlock()

	return statellite.UserReq[key]
}
