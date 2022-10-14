package star_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/zhangrt/voyager1_core/auth/grpc"
	"github.com/zhangrt/voyager1_core/auth/grpc/pb"
	"github.com/zhangrt/voyager1_core/auth/grpc/service"
	"github.com/zhangrt/voyager1_core/auth/luna"
	"github.com/zhangrt/voyager1_core/auth/star"
	"github.com/zhangrt/voyager1_core/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestStar(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	global.G_CONFIG.Zinx.Host = "127.0.0.1"
	global.G_CONFIG.Zinx.TcpPort = 2777
	go func() {
		star.StartClient()
	}()

	wg.Wait()
}

func TestGrpc(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	global.G_CONFIG.JWT.SigningKey = "gsafety"
	global.G_CONFIG.AUTHKey.RefreshToken = "new-token"
	global.G_CONFIG.System.UseMultipoint = true
	global.G_CONFIG.JWT.ExpiresTime = 10
	global.G_CONFIG.JWT.BufferTime = 7
	global.G_CONFIG.Grpc.Server.Host = "127.0.0.1"
	global.G_CONFIG.Grpc.Server.Port = 5431
	global.G_CONFIG.Grpc.Client.Host = "127.0.0.1"
	global.G_CONFIG.Grpc.Client.Port = 5431
	global.G_CONFIG.Grpc.Server.Network = "tcp"
	global.G_LOG = zap.New(zapcore.NewTee(), zap.AddCaller())
	luna.RegisterCasbin(&luna.UnimplementedCasbin{}) // 注入Casbin实现类
	luna.RegisterJwt(&luna.UnimplementedJwt{})       // 注入Jwt实现类

	go func() {
		defer wg.Done()
		grpc.NewServer().RegisterAuthServiceServer(&service.AuthService{}).LunchGrpcServer()
	}()
	time.Sleep(time.Second * 1)
	go func() {
		j := luna.NewTOKEN() // 唯一签名
		claims := j.CreateClaims(luna.BaseClaims{
			UUID:    uuid.NewV4(),
			ID:      100001,
			Name:    "test",
			Account: "test",
			RoleIds: []string{"101"},
		})
		token, _ := j.CreateToken(claims)

		for {
			time.Sleep(time.Second * 2)
			conn, client := star.GetGrpcClient()
			r, e := client.GrantedAuthority(context.Background(), &pb.Policy{
				Token: token,
			})
			if e != nil {
				fmt.Errorf(e.Error())
			}
			fmt.Println("Result:", r)
			fmt.Println()
			fmt.Println("----------------------------------------------")
			if r.NewToken != "" {
				token = r.NewToken
			}
			if r.Msg == "Authorization has expired" {
				break
			}
			star.CloseConn(conn)
		}
	}()
	wg.Wait()
}
