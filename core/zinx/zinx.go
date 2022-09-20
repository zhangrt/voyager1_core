package main

import (
	"github.com/aceld/zinx/znet"
)

func main() {
	s := znet.NewServer()
	s.Serve()
}
