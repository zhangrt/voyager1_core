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

type Server struct {
	srv pb.AuthServiceServer
}

func (server *Server) LuanGrpcServer() {

	fmt.Println(lunaLogo)

	s := grpc.NewServer()

	lis, _ := net.Listen(global.G_CONFIG.Grpc.Server.Network,
		fmt.Sprintf(global.G_CONFIG.Grpc.Server.Host, ":", global.G_CONFIG.Grpc.Server.Port))

	if server.srv != nil {
		pb.RegisterAuthServiceServer(s, server.srv)
	} else {
		pb.RegisterAuthServiceServer(s, new(service.AuthService))
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *Server) RegisterAuthServiceServer(as pb.AuthServiceServer) {
	s.srv = as
}
