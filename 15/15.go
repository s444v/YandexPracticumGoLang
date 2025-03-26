package main

import "fmt"

func GCD(a int, b int) int {
	if a <= 0 || b <= 0 {
		return 0
	}
	if a%b == 0 {
		return b
	}
	return GCD(b, a%b)
}

func main() {
	pairs := [][]int{
		{0, 5, 0},
		{7, -350, 0},
		{1460, 124, 4},
		{198, 42, 6},
		{1386, 4494, 42},
		{1470, 1575, 105},
	}
	for _, pair := range pairs {
		gcd := GCD(pair[0], pair[1])
		if gcd != pair[2] {
			fmt.Printf("НОД(%d, %d): %d != %d\n", pair[0], pair[1], gcd, pair[2])
		}
	}
	fmt.Println("Тестирование завершено")
}
