package main

import (
	"fmt"
	"sync"
	"time"
)

// количество горутин и ёмкость канала
const Count = 4

var wg sync.WaitGroup

// функция обработчик
func worker(ch chan int, i int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("ch = ", i, "v = ", v)
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	ch := make(chan int, Count)

	// создаём горутины
	for i := 1; i <= Count; i++ {
		wg.Add(1)
		go worker(ch, i)
	}

	// отправляем значения в канал
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
}
