package main

import (
	"fmt"
	"time"
)

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100000; j++ {
				sum++
			}
		}()
	}
	time.Sleep(2 * time.Millisecond)
	fmt.Println(sum)
}
