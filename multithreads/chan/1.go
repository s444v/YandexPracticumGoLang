package main

import (
	"fmt"
	"time"
)

var ch chan string

func hello() {
	time.Sleep(1500 * time.Millisecond)
	ch <- "Привет, гофер!"
}

func main() {
	ch = make(chan string)

	go hello()

	fmt.Println("Ждём строку для печати")
	s := <-ch
	fmt.Println(s)
}
