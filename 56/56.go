package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// используем в качестве сида time.Now().Unix(), чтобы при каждом запуске получался новый результат
	rnd := rand.NewSource(time.Now().Unix())
	data := make([]string, 0, 10000)
	for i := 0; i < 10000; i++ {
		// генерируем строку и добавляем в слайс
		data = append(data, generateString(10, rnd))
	}
	// засекаем время начала сортировки
	start := time.Now()
	sortStringSlice(data)
	// получем время, затраченное на сортировку
	finish := time.Since(start)

	// выводим результат
	fmt.Printf("сортировка выполнялась %f секунд\n", finish.Seconds())
}

// sortStringSlice сортирует слайс строк
func sortStringSlice(in []string) {
	sort.Slice(in, func(i, j int) bool {
		return in[i] < in[j]
	})
}

// generateString генерирует случайную строку длины n из английских прописных букв
func generateString(n int, generator rand.Source) string {
	result := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		// генерируем случайное число
		randomNumber := generator.Int63()
		// английские буквы лежат в диапазоне от 97 до 122, поэтому:
		// 1) берем остаток от деления случайного числа на 26, получая диапазон [0,25]
		// 2) прибавляем к полученному числу 97 и получаем итоговый интервал: [97, 122].
		result = append(result, byte(randomNumber%26+97))
	}
	return string(result)
}
