package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		go func(v int) {
			wg.Add(1)
			fmt.Print(v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
