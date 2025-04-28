package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func asterisks(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			// выведет context deadline exceeded
			fmt.Println(ctx.Err())
			return
		case <-time.After(400 * time.Millisecond):
			fmt.Print("*")
		}
	}
}

func dots(ctx context.Context) {
	defer wg.Done()
	// указываем ctx в качестве родительского контекста
	ctxAsterisk, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()
	go asterisks(ctxAsterisk)
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(300 * time.Millisecond):
			fmt.Print(".")
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	wg.Add(2)
	go dots(ctx)
	wg.Wait()
}
