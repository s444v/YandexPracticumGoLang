package main

import (
	"fmt"
	"net/http"
)

func main() {
	response, err := http.Get("http://localhost:8080/time")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Статус ответа:", response.StatusCode)
}
