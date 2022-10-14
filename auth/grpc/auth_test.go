package grpc_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"runtime"
	"sync"
	"testing"

	uuid "github.com/satori/go.uuid"
	pb "github.com/zhangrt/voyager1_core/auth/grpc/pb"
	service "github.com/zhangrt/voyager1_core/auth/grpc/service"
	luna "github.com/zhangrt/voyager1_core/auth/luna"

	"github.com/zhangrt/voyager1_core/global"
	"google.golang.org/grpc"
)

func TestAuth(t *testing.T) {
	// 测试所需的环境配置参数
	global.G_CONFIG.JWT.SigningKey = "gsafety"
	global.G_CONFIG.AUTHKey.RefreshToken = "new-token"
	global.G_CONFIG.System.UseMultipoint = true
	global.G_CONFIG.JWT.ExpiresTime = 60
	// global.BlackCache = local_cache.NewCache(
	// 	local_cache.SetDefaultExpire(time.Second * time.Duration(1800)),
	// )
	// global.BlackCache.SetDefault("", struct{}{})
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(3)

	{
		go func() {
			defer wg.Done()

			s := grpc.NewServer()

			lis, _ := net.Listen("tcp", "127.0.0.1:8081")

			pb.RegisterAuthServiceServer(s, new(service.AuthService))
			if err := s.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}

		}()
	}

	{
		go func() {
			// time.Sleep(time.Second * 3)

			for {
				conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
				if err != nil {
					log.Fatal(err)
				}
				authClient := pb.NewAuthServiceClient(conn)
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
				res, err := authClient.GetUser(context.Background(), &pb.Token{
					Token: token,
				})
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(res)
				conn.Close()

				// time.Sleep(time.Second * 5)
			}

		}()
	}
	wg.Wait()
}
