package main

import (
	"fmt"
	"time"
)

func main() {
	m := make([]int, 100)
	for i := range len(m) {
		go func() {
			m[i] = 1
		}()
	}
	time.Sleep(200 * time.Millisecond)
	sum := 0
	for i := 0; i < len(m); i++ {
		sum += m[i]
	}
	fmt.Println(sum)
}
