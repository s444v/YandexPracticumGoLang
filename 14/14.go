package main

import (
	"fmt"
	"math"
)

var Message = `Добро пожаловать в самую лучшую программу для вычисления 
квадратного корня из заданного числа`

func CalculateSquareRoot(my_number float64) float64 {
	return math.Sqrt(my_number)
}

func Calc(number float64) {
	if number < 0 {
		return
	} else {
		fmt.Println("Мы вычислили квадратный корень из введённого вами числа. Это будет:", CalculateSquareRoot(number))
	}
}

func main() {
	fmt.Println(Message)
	Calc(25.5)
}
