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
	global.G_CONFIG.JWT.ExpiresTime = 60
	global.G_CONFIG.Grpc.Server.Host = "127.0.0.1"
	global.G_CONFIG.Grpc.Server.Port = 5431
	global.G_CONFIG.Grpc.Client.Host = "127.0.0.1"
	global.G_CONFIG.Grpc.Client.Port = 5431
	global.G_CONFIG.Grpc.Server.Network = "tcp"

	go func() {
		defer wg.Done()
		grpc.NewServer().RegisterAuthServiceServer(&service.AuthService{}).LunchGrpcServer()
	}()
	time.Sleep(time.Second * 3)
	go func() {
		for {
			j := luna.NewTOKEN() // 唯一签名
			claims := j.CreateClaims(luna.BaseClaims{
				UUID:    uuid.NewV4(),
				ID:      100001,
				Name:    "test",
				Account: "test",
				RoleIds: []string{"101"},
			})
			token, err := j.CreateToken(claims)
			if err != nil {
				continue
			}
			conn, client := star.GetGrpcClient()
			r, e := client.GetUser(context.Background(), &pb.Token{
				Token: token,
			})
			if e != nil {
				fmt.Errorf(e.Error())
			}
			fmt.Println(r)
			star.CloseConn(conn)
			// time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
