package star_test

import (
	"sync"
	"testing"

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
