package grpc_test

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/zhangrt/voyager1_core/auth/grpc"
	"github.com/zhangrt/voyager1_core/auth/grpc/service"
	"github.com/zhangrt/voyager1_core/auth/luna"
	"github.com/zhangrt/voyager1_core/auth/star"
	pb "github.com/zhangrt/voyager1_core/com/gs/voyager1_core/auth/grpc/pb"
	"github.com/zhangrt/voyager1_core/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestGrpcJ(t *testing.T) {
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
		grpc.NewServerJ().RegisterAuthServiceServer(&service.AuthServiceJ{}).LunchGrpcServerJ()
	}()
	time.Sleep(time.Second * 1)
	go func() {
		j := luna.NewTOKEN() // 唯一签名
		claims := j.CreateClaims(luna.BaseClaims{
			ID:      strings.ReplaceAll(uuid.NewV4().String(), "-", ""),
			Name:    "test",
			Account: "test",
			RoleIds: []string{"101"},
		})
		token, _ := j.CreateToken(claims)

		for {
			time.Sleep(time.Second * 2)
			conn, client := star.GetGrpcClientJ()
			defer star.CloseConn(conn)
			r, e := client.GrantedAuthority(context.Background(), &pb.Policy{
				Token: token,
			})
			if e != nil {
				fmt.Errorf(e.Error())
			}
			fmt.Println("Result:", r)
			fmt.Println()
			if r.NewToken != "" {
				token = r.NewToken
			}
			if r.Msg == "Authorization has expired" {
				break
			}
			u, err := client.GetUser(context.Background(), &pb.Authentication{
				Token: token,
			})
			if err == nil {
				fmt.Println("USER:", u)
			}
			fmt.Println("----------------------------------------------")
		}
	}()
	wg.Wait()
}
