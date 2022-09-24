package main

import (
	"runtime"
	"sync"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/zhangrt/voyager1_core/constant"
	"github.com/zhangrt/voyager1_core/global"
	"github.com/zhangrt/voyager1_core/zinx/api"
	"github.com/zhangrt/voyager1_core/zinx/core/luna"
	"github.com/zhangrt/voyager1_core/zinx/core/star"
	pb "github.com/zhangrt/voyager1_core/zinx/pb"
)

func TestSer(t *testing.T) {
	// 测试所需的环境配置参数
	global.G_CONFIG.JWT.SigningKey = "gsafety"
	global.G_CONFIG.AUTHKey.RefreshToken = "new-token"
	global.G_CONFIG.System.UseMultipoint = true
	global.G_CONFIG.JWT.ExpiresTime = 60
	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Second * time.Duration(1800)),
	)
	global.BlackCache.SetDefault("", struct{}{})

	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(2)

	{
		s := luna.Server(
			luna.Router{
				ID:     constant.TOKEN_REQ,
				ROUTER: &api.AuthorizationRequestApi{},
			},
			luna.Router{
				ID:     constant.POLICY_REQ,
				ROUTER: &api.AuthenticationRequestApi{},
			},
			luna.Router{
				ID:     constant.USER_REQ,
				ROUTER: &api.UserRequestApi{},
			},
			luna.Router{
				ID:     constant.HEARTBEAT_REQ,
				ROUTER: &api.HeartbeatRequestApi{},
			},
		)

		go func() {
			defer wg.Done()
			go s.Serve()
		}()

	}

	{
		client1 := star.NewTcpClient("127.0.0.1", 8999)
		client2 := star.NewTcpClient("127.0.0.1", 8999)

		go func() {
			star.StatelliteMgrObj.ClientObj[0] = client1
			go client1.Start()

		}()

		go func() {
			star.StatelliteMgrObj.ClientObj[1] = client2
			go client2.Start()

		}()

		go func() {
			for {
				token := pb.Token{
					Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZjBjY2M1ZGEtMzA5NC00MWJkLWJmN2UtNzE1MDZjMTdjNDQ3IiwiSUQiOjc5NDI4MDIxMTY5MjgxNDMzNywiQWNjb3VudCI6InRlc3QiLCJOYW1lIjoiQklHIE1vbnN0ZXIiLCJBdXRob3JpdHlJZCI6Ijk1MjgiLCJBdXRob3JpdHkiOnsiQ3JlYXRlZEF0IjoiMjAyMi0wOS0wNlQxOTo1ODowMy40MTM1MDgrMDg6MDAiLCJVcGRhdGVkQXQiOiIyMDIyLTA5LTA2VDE5OjU4OjA0LjY1NDI4MSswODowMCIsIkRlbGV0ZWRBdCI6bnVsbCwiYXV0aG9yaXR5SWQiOiI5NTI4IiwiYXV0aG9yaXR5TmFtZSI6Iua1i-ivleinkuiJsiIsInBhcmVudElkIjoiMCIsImRlZmF1bHRSb3V0ZXIiOiI0MDQifSwiQXV0aG9yaXRpZXMiOm51bGwsIkRlcGFydE1lbnRJZCI6IiIsIkRlcGFydE1lbnROYW1lIjoiIiwiVW5pdElkIjoiIiwiVW5pdE5hbWUiOiIiLCJCdWZmZXJUaW1lIjoxMjAsImV4cCI6MTY2MzkwODYxMiwiaXNzIjoiZ3NhZmV0eSIsIm5iZiI6MTY2MzkwNzQzMn0.q3r3QwpLGcAq45OHinhB1wncEbATCjXwKdbMApgXLVM",
				}

				uid := uuid.NewV4().String()

				// 请求
				star.StatelliteMgrObj.AddTokenReq(uid, &token)

				star.StatelliteMgrObj.AddMsgKey(uid, constant.TOKEN_REQ)

				for {
					// 结果
					r := star.StatelliteMgrObj.GetTokenResult(uid)
					if r != nil {
						println("Key >>>>> ", uid)
						println("<Clinet1111111111111() GetTokenResult Key>: ", r.Key)
						println("====================== %d =========================", uid == r.Key)
						star.StatelliteMgrObj.RemoveTokenResult(uid)
						break
					}
					// time.Sleep(time.Millisecond)
				}

				time.Sleep(time.Millisecond * 100)
			}
		}()

		go func() {
			for {
				token := pb.Token{
					Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZjBjY2M1ZGEtMzA5NC00MWJkLWJmN2UtNzE1MDZjMTdjNDQ3IiwiSUQiOjc5NDI4MDIxMTY5MjgxNDMzNywiQWNjb3VudCI6InRlc3QiLCJOYW1lIjoiQklHIE1vbnN0ZXIiLCJBdXRob3JpdHlJZCI6Ijk1MjgiLCJBdXRob3JpdHkiOnsiQ3JlYXRlZEF0IjoiMjAyMi0wOS0wNlQxOTo1ODowMy40MTM1MDgrMDg6MDAiLCJVcGRhdGVkQXQiOiIyMDIyLTA5LTA2VDE5OjU4OjA0LjY1NDI4MSswODowMCIsIkRlbGV0ZWRBdCI6bnVsbCwiYXV0aG9yaXR5SWQiOiI5NTI4IiwiYXV0aG9yaXR5TmFtZSI6Iua1i-ivleinkuiJsiIsInBhcmVudElkIjoiMCIsImRlZmF1bHRSb3V0ZXIiOiI0MDQifSwiQXV0aG9yaXRpZXMiOm51bGwsIkRlcGFydE1lbnRJZCI6IiIsIkRlcGFydE1lbnROYW1lIjoiIiwiVW5pdElkIjoiIiwiVW5pdE5hbWUiOiIiLCJCdWZmZXJUaW1lIjoxMjAsImV4cCI6MTY2MzkwODYxMiwiaXNzIjoiZ3NhZmV0eSIsIm5iZiI6MTY2MzkwNzQzMn0.q3r3QwpLGcAq45OHinhB1wncEbATCjXwKdbMApgXLVM",
				}

				uid := uuid.NewV4().String()

				// 请求
				star.StatelliteMgrObj.AddTokenReq(uid, &token)

				star.StatelliteMgrObj.AddMsgKey(uid, constant.TOKEN_REQ)

				for {
					// 结果
					r := star.StatelliteMgrObj.GetTokenResult(uid)
					if r != nil {
						println("Key <<<<<< ", uid)
						println("<Clinet2222222222222() GetTokenResult Key>: ", r.Key)
						println("====================== %d =========================", uid == r.Key)
						star.StatelliteMgrObj.RemoveTokenResult(uid)
						break
					}
					// time.Sleep(time.Millisecond)
				}

				time.Sleep(time.Millisecond * 200)
			}
		}()
	}

	wg.Wait()
}
