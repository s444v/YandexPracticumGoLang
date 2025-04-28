package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// создаём таймер, который сработает через одну секунду
	timer := time.NewTimer(10 * time.Second)
	// ожидаем срабатывания таймера
	t := <-timer.C
	// выводим разницу во времени, которая должна быть 1
	fmt.Println(t.Sub(start).Seconds())
}
