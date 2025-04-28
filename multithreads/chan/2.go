package main

import "fmt"

func main() {
	ch := make(chan string)
	go func(ch chan string) {
		fmt.Println(<-ch)
	}(ch)
	ch <- "Hello"
}
