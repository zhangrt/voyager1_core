package zinx

import (
	"fmt"
	"runtime"
	"sync"

	c "github.com/zhangrt/voyager1_core/zinx/core/client"
)

func client() {
	// 开启一个waitgroup，同时运行3个goroutine

	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			client := c.NewTcpClient("127.0.0.1", 8999)
			client.Start()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			client := c.NewTcpClient("127.0.0.1", 8999)
			client.Start()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			client := c.NewTcpClient("127.0.0.1", 8999)
			client.Start()
		}
	}()

	fmt.Println("client start")
	wg.Wait()
	fmt.Println("client exit")
}
