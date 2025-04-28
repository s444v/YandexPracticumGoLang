package main

import (
	"fmt"
	"sync"
)

var list []int
var wg sync.WaitGroup
var mu sync.Mutex

func do() {
	for i := 0; i < 100; i++ {
		mu.Lock()
		list = append(list, i)
		mu.Unlock()
	}
	wg.Done()
}

func main() {
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go do()
	}
	wg.Wait()
	// проверка содержимого у слайса
	sum := 0
	for _, v := range list {
		sum += v
	}
	fmt.Println(len(list), sum)
}
