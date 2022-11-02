package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/zhangrt/voyager1_core/auth/grpc/service"
	pb "github.com/zhangrt/voyager1_core/com/gs/voyager1_core/auth/grpc/pb"
	"github.com/zhangrt/voyager1_core/global"
	"google.golang.org/grpc"
)

var lunaLogoJ = `
 ██
 ██
 ██         ██      ██   ██▄████▄      ▄▄▀▀▀▀▄▄▄
 ██         ██      ██   ██▀   ██    █▀         ▀█
 ██         ██      ██   ██    ██    █           █
 ██▄▄▄▄▄▄   ██      ██   ██    ██    █         ▄▄█
 ▀▀▀▀▀▀▀▀   ▀▀▀▀▀▀▀▀ ▀   ▀▀    ▀▀    ▀▄▄▄▄▄▀▀▀▀  ▀▄▄
                                                        `
var topLineJ = `┌───────────────────────────────────────────────────┐`
var borderLineJ = `│`
var bottomLineJ = `└───────────────────────────────────────────────────┘`

type ServerJ struct {
	srv pb.AuthServiceServer
}

func NewServerJ() *ServerJ {
	return &ServerJ{}
}

func (server *ServerJ) LunchGrpcServerJ() {

	printLogoJ()

	s := grpc.NewServer()

	lis, _ := net.Listen(global.G_CONFIG.Grpc.Server.Network,
		fmt.Sprintf("%s:%d", global.G_CONFIG.Grpc.Server.Host, global.G_CONFIG.Grpc.Server.Port))

	if server.srv != nil {
		pb.RegisterAuthServiceServer(s, server.srv)
	} else {
		pb.RegisterAuthServiceServer(s, new(service.AuthServiceJ))
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *ServerJ) RegisterAuthServiceServer(as pb.AuthServiceServer) *ServerJ {
	s.srv = as
	return s
}

func printLogoJ() {
	fmt.Println(lunaLogoJ)
	fmt.Println(topLineJ)
	v := fmt.Sprintf("%s version:v0.1                                      %s", borderLine, borderLine)
	e := fmt.Sprintf("%s email:zhoujiajun@gsafety.com                      %s", borderLine, borderLine)
	fmt.Println(v)
	fmt.Println(e)
	fmt.Println(bottomLineJ)

	fmt.Printf("[Luna] started at %s, by %s . ",
		fmt.Sprintf("%s:%d", global.G_CONFIG.Grpc.Server.Host, global.G_CONFIG.Grpc.Server.Port),
		fmt.Sprintf(global.G_CONFIG.Grpc.Server.Network))
	fmt.Println()
}
