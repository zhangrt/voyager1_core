package luna

import (
	"sync"
)

func Luna() {
	var wg sync.WaitGroup
	wg.Add(2)
	s := NewLuna()
	go func() {
		defer wg.Done()
		go s.Serve()
	}()
}
