package main

// ##CGO_ENABLED=0
import "C"

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	counter := 0
	for i := 0; i < 5; i++ {
		go func() {
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("OK", counter)
}
