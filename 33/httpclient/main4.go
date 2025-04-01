package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	for {
		response, err := http.Get("http://localhost:8080/time")
		if err != nil {
			fmt.Println(err)
			break
		}
		// читаем тело ответа
		body, err := io.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			fmt.Println("Ошибка чтения:", err)
			break
		}
		fmt.Println(string(body))
		// ждём 5 секунд
		time.Sleep(5 * time.Second)
	}
}
