package main

import "fmt"

func avg(x, y int) int {
	return x + y/2
}

func main() {
	fmt.Println(avg(5, 10))
}

func div(x, y int) (int, error) {
	// вставьте проверку y на 0 и добавьте возвращение ошибки
	// div(x, y int) (int, error)
	if y == 0 {
		return 0, nil
	}
	return x / y, nil
}
