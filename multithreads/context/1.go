package main

import (
	"context"
	"fmt"
	"time"
)

func process(ctx context.Context) {
	fmt.Println("Начало работы")
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Прервали работу")
			return
		default:
			fmt.Println(i)
			time.Sleep(500 * time.Millisecond)
		}
	}
	fmt.Println("Конец работы")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// через 2 секунды вызываем cancel для отмены операции
	time.AfterFunc(2*time.Second, cancel)

	process(ctx)
}
