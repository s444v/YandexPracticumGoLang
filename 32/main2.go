package main

import (
	"fmt"
	"net/http"
)

func mainHandle(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Получен запрос")
}

func main() {
	fmt.Println("Запускаем сервер")
	http.HandleFunc(`/`, mainHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Завершаем работу")
}
