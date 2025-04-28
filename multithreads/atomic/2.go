package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var counter atomic.Uint64

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.Add(1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Load())
}
