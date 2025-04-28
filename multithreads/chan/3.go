package main

import (
	"fmt"
)

func do(in, out chan int) {
	defer close(out) // закрываем канал по окончании работы функции
	for v := range in {
		// тело цикла
		// отправляем результат в другой канал
		out <- 2 * v
	}
}

func main() {
	chIn := make(chan int)
	chOut := make(chan int)

	// запускаем горутину, которая преобразует числа
	go do(chIn, chOut)

	// горутина, которая отправляет числа на обработку
	go func() {
		// закрываем канал после отправки всех чисел
		defer close(chIn)

		for i := 0; i <= 50; i++ {
			chIn <- i
		}
	}()
	// в цикле читаем числа из результирующего канала
	for v := range chOut {
		// тело цикла

		fmt.Print(v, " ")
	}
}
