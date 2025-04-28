package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1000 * time.Millisecond)
	timer := time.NewTimer(5 * time.Second)

	var count int
	for {
		select {
		case <-ticker.C:
			count++
			fmt.Print(count, " ")
		case <-timer.C:
			fmt.Println("Расчёт закончен")
			return
		}
	}
}
