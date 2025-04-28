package main

import (
	"fmt"
)

func main() {
	напишите код программы
	ticker := time.NewTicker(200 * time.Millisecond)

	for i := 1; i <= 20; i++ {
		select {
		case <-ticker.C:
			if i%5 == 0 {
				fmt.Print("o")

				continue
			}
			fmt.Print(".")
		}
	}
	//fmt.Println(factorial(4))
}
// func factorial(n int) int {
// 	// sum := 1
// 	// for i := 1; i <= n; i++ {
// 	// 	sum *= i
// 	// }
// 	// return sum
// 	if n == 1 {
// 		return n
// 	}
// 	return n * factorial(n-1)
// }


