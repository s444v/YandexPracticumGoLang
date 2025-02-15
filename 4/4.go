package main

import (
	"fmt"
	"time"
)

func whatTime(city string) time.Time {
	var offsetUTC = map[string]int{
		"Санкт-Петербург": 3,
		"Москва":          3,
		"Самара":          4,
		"Новосибирск":     7,
		"Екатеринбург":    5,
		"Казань":          3,
		"Омск":            6,
		"Хабаровск":       10,
		"Пермь":           5,
		"Краснодар":       3,
		"Калининград":     2,
	}
	// напишите код, который берёт текущее UTC-время
	// и добавляет к нему сдвиг города в часах
	add := time.Duration(offsetUTC[city])
	return (time.Now().UTC()).Add(add * time.Hour)
}

func main() {
	for _, city := range []string{"Екатеринбург", "Москва", "Хабаровск"} {
		t := whatTime(city)
		fmt.Printf("%s: %d ч.\n", city, t.Hour())
	}
}
