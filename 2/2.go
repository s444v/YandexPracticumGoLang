package main

import (
	"fmt"
	"time"
)

func main() {
	// дата выхода первой серии
	begin := time.Date(2011, 4, 17, 0, 0, 0, 0, time.UTC)
	// укажите дату выхода последней серии
	end := time.Date(2019, 4, 15, 0, 0, 0, 0, time.UTC)
	// вычислите, сколько времени шёл сериал
	duration := end.Sub(begin)
	// разделите количество часов duration.Hours() на 24,
	// чтобы получить количество дней
	fmt.Printf("Сериал шёл %d дней", int(duration.Hours())/24)
}
