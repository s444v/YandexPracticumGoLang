package main

import (
	"fmt"
	"time"
)

func whatTime(friend string) time.Time {
	var friends = map[string]string{
		"Серёга": "Омск",
		"Соня":   "Москва",
		"Дима":   "Екатеринбург",
		"Алина":  "Новосибирск",
		"Егор":   "Калининград",
	}
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

	// напишите недостающий код
	add := time.Duration(offsetUTC[friends[friend]])
	return (time.Now().UTC()).Add(add * time.Hour)
}

func main() {
	for _, friend := range []string{"Соня", "Дима", "Егор"} {
		t := whatTime(friend)
		fmt.Printf("%s: %d ч.\n", friend, t.Hour())
	}
}
