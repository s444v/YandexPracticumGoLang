package main

import (
	"fmt"
	"sync"
	"time"
)

// можно определить переменную в main() перед циклом
var mu sync.Mutex

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		go func() {
			mu.Lock()
			for j := 0; j < 100000; j++ {
				sum++
			}
			mu.Unlock()
		}()
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println(sum)
}
