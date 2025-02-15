package main

import (
	"fmt"
	"time"
)

func main() {
	// создайте переменные типа time.Time
	start := time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2023, 8, 8, 0, 0, 0, 0, time.UTC)
	// получите интервал между датами
	// вычислите количество дней
	days := end.Sub(start).Hours() / 24
	fmt.Printf("Между выходами версий прошло %d дней", int(days))
}
