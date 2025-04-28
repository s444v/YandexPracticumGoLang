package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu sync.RWMutex
	m  map[int]int
}

func (cache *Cache) Get(i int) int {
	cache.mu.RLock()
	v, ok := cache.m[i]
	if ok {
		cache.mu.RUnlock()
		return v
	}
	cache.mu.RUnlock()
	// получаем значение для указанного ключа
	cache.mu.Lock()
	defer cache.mu.Unlock()
	v = 2 * i
	cache.m[i] = v
	return v
}

func main() {
	cache := Cache{
		m: make(map[int]int),
	}
	for i := 0; i < 20; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				cache.Get(j)
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(len(cache.m))
}
