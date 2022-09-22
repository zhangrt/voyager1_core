package main

import (
	"runtime"
	"sync"

	"github.com/zhangrt/voyager1_core/zinx/api"
	"github.com/zhangrt/voyager1_core/zinx/core"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(2)

	s := core.Server(
		core.Router{
			ID:     1,
			ROUTER: &api.AuthorizationApi{},
		},
		core.Router{
			ID:     2,
			ROUTER: &api.AuthenticationRequestApi{},
		},
	)

	go func() {
		defer wg.Done()
		go s.Serve()
	}()

	client := core.NewTcpClient("127.0.0.1", 8999)
	go func() {
		go client.Start()
	}()

	wg.Wait()

}
