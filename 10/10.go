package main

import (
	"fmt"
	"time"
)

func Counter(count int, t time.Time) string {
	date := t.Format("02.01.2006")
	return fmt.Sprintf("%d %s", count, date)
}

func main() {
	now := time.Now()
	for count := 1; count < 4; count++ {
		fmt.Println(Counter(count, now))
	}
}
