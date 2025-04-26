package main

import "fmt"

func main() {
	// Исходный массив
	numbers := [...]int{1, 3, 8, 19, 42}

	// 1. Создаём слайс указателей на элементы массива
	pointers := make([]*int, len(numbers))
	for i := range numbers {
		pointers[i] = &numbers[i]
	}

	// 2. Умножаем на 3 значения, на которые ссылаются указатели
	for _, ptr := range pointers {
		*ptr *= 3
	}

	// Выводим изменённый массив
	fmt.Println(numbers) // [3 9 24 57 126]
}
