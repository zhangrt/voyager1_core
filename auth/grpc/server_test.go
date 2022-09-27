package grpc_test

import (
	"net"
	"testing"

	"google.golang.org/grpc"
)

func TestServer(t *testing.T) {
	rpcServer := grpc.NewServer()

	l, _ := net.Listen("tcp", "8081")

	rpcServer.Serve(l)
}
