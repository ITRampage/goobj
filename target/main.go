package main

import (
	"sync"
	"time"

	"github.com/ITRampage/goobj/inc"
)

func main() {
	wg := &sync.WaitGroup{}
	o := inc.New(0)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go (func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				o.Inc()
				if j%1000 == 0 {
					o.Print()
				}
			}
		})()
	}
	wg.Wait()
	o.Descruct()
	time.Sleep(1 * time.Second)
}
