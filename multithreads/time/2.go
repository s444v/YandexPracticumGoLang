package main

import (
	"fmt"
	"time"
)

var quit = make(chan struct{})

func hello() {
	fmt.Println("Упс!")
	quit <- struct{}{}
}

func main() {
	fmt.Println("Сейчас вылетит птичка")
	time.AfterFunc(2*time.Second, hello)
	for i := 1; i <= 3; i++ {
		v := i
		time.AfterFunc(time.Duration(500*i)*time.Millisecond, func() {
			fmt.Println(v)
		})
	}
	<-quit
}
