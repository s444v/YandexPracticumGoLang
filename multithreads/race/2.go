package main

import (
	"fmt"
	"sync"
	"time"
)

type Map struct {
	mu sync.RWMutex
	m  map[string]string
}

func (m *Map) Get(key string) string {
	// RLock() даёт возможность нескольким горутинам читать данные
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.m[key]
}

func (m *Map) Set(key string, v string) {
	// Lock() блокирует все операции
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[key] = v
}

func main() {
	m := Map{
		m: make(map[string]string),
	}
	for i := 0; i < 10; i++ {
		go func() {
			for {
				m.Set("a", ".")
				time.Sleep(50 * time.Millisecond)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			count := 0
			for {
				v := m.Get("a")
				count++
				if count%10 == 0 {
					fmt.Print(v)
				}
				time.Sleep(20 * time.Millisecond)
			}
		}()
	}
	time.Sleep(1 * time.Second)
}
