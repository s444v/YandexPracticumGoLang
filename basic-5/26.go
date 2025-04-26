package main

import "fmt"

// вставьте сюда версию DivSizes без указателей
func DivSizes(x, y int) (int, int) {
	x /= 2
	y /= 2
	return x, y
}

func main() {
	x, y := DivSizes(10, 20)
	fmt.Println(x, y)
}
