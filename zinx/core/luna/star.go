package luna

import (
	"fmt"
	"sync"
	"time"

	"github.com/aceld/zinx/ziface"
	auth "github.com/zhangrt/voyager1_core/auth/luna"
	"github.com/zhangrt/voyager1_core/constant"
	"github.com/zhangrt/voyager1_core/global"
	pb "github.com/zhangrt/voyager1_core/zinx/pb"
	"google.golang.org/protobuf/proto"
)

// 客户端 star 由 Server 维护管理的一组 client
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

}

var jwt = auth.NewJWT()

// 验证token合法并将结果发送回客户端 star
func (s *Star) CheckToken(req *pb.Token) {
	token := req.Token
	msg := &pb.Result{}
	msg.Key = req.Key

	if jwt.IsBlacklist(token) {
		msg.Success = false
		msg.Msg = "Your account is off-site logged in or the token is invalid"
		s.SendMsg(constant.TOKEN_RES, msg)
		return
	}

	j := auth.NewTOKEN()
	claims, err := j.ParseToken(token)
	if err != nil {
		if err == auth.TokenExpired {
			msg.Success = false
			msg.Msg = "Authorization has expired"
		} else {
			msg.Success = false
			msg.Msg = err.Error()
		}
		s.SendMsg(constant.TOKEN_RES, msg)
		return
	}

	now := time.Now().Unix()
	if claims.ExpiresAt-now < claims.BufferTime {
		claims.ExpiresAt = now + global.G_CONFIG.JWT.ExpiresTime
		newToken, _ := j.CreateTokenByOldToken(token, *claims)
		newClaims, _ := j.ParseToken(newToken)

		// 单点登录
		if !global.G_CONFIG.System.UseMultipoint {
			// 获取缓存中account的未过期token
			RedisJwtToken, err := jwt.GetCacheJWT(newClaims.Account)
			if err != nil {
				msg.Success = false
				msg.Msg = "get cache jwt failed"
			} else {
				// 当之前的取成功时才进行拉黑操作
				// refresh token
				msg.Success = true
				msg.Msg = global.G_CONFIG.AUTHKey.RefreshToken + constant.MARKER + newToken
				msg.Claims = protoTransformClaims(newClaims)
				_ = jwt.JsonInBlacklist(auth.JwtBlacklist{Jwt: RedisJwtToken})
			}
			// 无论如何都要记录当前的活跃状态
			_ = jwt.SetCacheJWT(newToken, newClaims.Account)

		} else {
			// 启用多点登录
			msg.Success = true
			msg.Msg = global.G_CONFIG.AUTHKey.RefreshToken + constant.MARKER + newToken
			msg.Claims = protoTransformClaims(newClaims)
		}
		s.SendMsg(constant.TOKEN_RES, msg)
		return
	}
	msg.Msg = "authorization success"
	msg.Success = true
	msg.Claims = protoTransformClaims(claims)
	s.SendMsg(constant.TOKEN_RES, msg)
}

// 验证角色权限
func (s *Star) AuthenticationRequest(req *pb.Policy) {
	msg := &pb.Result{}
	msg.Key = req.Key
	success, _ := auth.CheckPolicy(req.AuthorityId, req.Path, req.Method)
	msg.Success = success
	if !success {
		msg.Msg = "insufficient privileges"
	}
	s.SendMsg(constant.POLICY_RES, msg)
}

// 获取用户信息
func (s *Star) GetUserInfo(req *pb.Token) {
	msg := &pb.User{}
	msg.Key = req.Key
	claims, _ := auth.GetUser(req.Token)
	if claims != nil {
		msg.Claims = protoTransformClaims(claims)
		msg.UserID = int64(claims.ID)
		msg.UUID = claims.UUID.String()
		msg.AuthorityId = claims.RoleId
	}
	s.SendMsg(constant.USER_RES, msg)
}

func (s *Star) Receipe(id int32) {
	msg := &pb.Receipe{}
	msg.Id = s.PID
	s.SendMsg(constant.HEARTBEAT_RES, msg)
}

// 转换
func protoTransformClaims(c *auth.CustomClaims) *pb.CustomClaims {
	p := pb.CustomClaims{
		Claims: &pb.BaseClaims{
			UserID:      int64(c.BaseClaims.ID),
			UUID:        c.BaseClaims.UUID.String(),
			AuthorityId: c.BaseClaims.RoleId,
			Account:     c.BaseClaims.Account,
			Name:        c.BaseClaims.Name,
		},
		BufferTime: c.BufferTime,
		Standard: &pb.StandardClaims{
			Audience:  c.StandardClaims.Audience,
			ExpiresAt: c.StandardClaims.ExpiresAt,
			Id:        c.StandardClaims.Id,
			IssuedAt:  c.StandardClaims.IssuedAt,
			Issuer:    c.StandardClaims.Issuer,
			NotBefore: c.StandardClaims.NotBefore,
			Subject:   c.StandardClaims.Subject,
		},
	}
	return &p
}
