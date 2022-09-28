package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/zhangrt/voyager1_core/auth/grpc/pb"
	"github.com/zhangrt/voyager1_core/auth/grpc/service"
	"github.com/zhangrt/voyager1_core/global"
	"google.golang.org/grpc"
)

var lunaLogo = `                                        
 ██                        
 ██                        
 ██         ██      ██   ██▄████▄      ▄▄▀▀▀▀▄▄▄
 ██         ██      ██   ██▀   ██    █▀         ▀█
 ██         ██      ██   ██    ██    █           █
 ██▄▄▄▄▄▄   ██      ██   ██    ██    █         ▄▄█
 ▀▀▀▀▀▀▀▀   ▀▀▀▀▀▀▀▀ ▀   ▀▀    ▀▀    ▀▄▄▄▄▄▀▀▀▀  ▀▄▄
                                                        `
var topLine = `┌───────────────────────────────────────────────────┐`
var borderLine = `│`
var bottomLine = `└───────────────────────────────────────────────────┘`

type Server struct {
	srv pb.AuthServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (server *Server) LunchGrpcServer() {

	printLogo()

	s := grpc.NewServer()

	lis, _ := net.Listen(global.G_CONFIG.Grpc.Server.Network,
		fmt.Sprintf("%s:%d", global.G_CONFIG.Grpc.Server.Host, global.G_CONFIG.Grpc.Server.Port))

	if server.srv != nil {
		pb.RegisterAuthServiceServer(s, server.srv)
	} else {
		pb.RegisterAuthServiceServer(s, new(service.AuthService))
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *Server) RegisterAuthServiceServer(as pb.AuthServiceServer) *Server {
	s.srv = as
	return s
}

func printLogo() {
	fmt.Println(lunaLogo)
	fmt.Println(topLine)
	v := fmt.Sprintf("%s version:v0.1                                      %s", borderLine, borderLine)
	e := fmt.Sprintf("%s email:zhoujiajun@gsafety.com                      %s", borderLine, borderLine)
	fmt.Println(v)
	fmt.Println(e)
	fmt.Println(bottomLine)

	fmt.Printf("[Luna] started at %s, by %s . ",
		fmt.Sprintf("%s:%d", global.G_CONFIG.Grpc.Server.Host, global.G_CONFIG.Grpc.Server.Port),
		fmt.Sprintf(global.G_CONFIG.Grpc.Server.Network))
	fmt.Println()
}
