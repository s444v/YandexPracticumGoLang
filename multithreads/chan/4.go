package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			ch1 <- 0
			time.Sleep(50 * time.Millisecond)
		}
	}()
	go func() {
		for {
			ch2 <- 1
			time.Sleep(50 * time.Millisecond)
		}
	}()
	var v int
	for i := 0; i < 50; i++ {
		select {
		case v = <-ch1:
			fmt.Print(v, " ")
		case v = <-ch2:
			fmt.Print(v, " ")
		}
	}
}
