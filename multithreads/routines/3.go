package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	const n = 3

	for i := 0; i < n; i++ {
		wg.Add(1) // инкрементируем счётчик перед запуском горутины

		go func(i int) {
			// уменьшаем счётчик, когда горутина завершает работу
			defer wg.Done()

			for j := 0; j < 10; j++ {
				fmt.Print(i)
			}
		}(i)
	}

	wg.Wait() // ждём завершения всех горутин
	fmt.Println("\nконец")
}
