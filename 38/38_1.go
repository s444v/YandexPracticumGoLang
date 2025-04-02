package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.PostForm("http://localhost:8080/", url.Values{})
	if err != nil {
		fmt.Println("Ошибка", err)
		return
	}
	fmt.Println("Код статуса:", resp.StatusCode)
	// читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("Ошибка", err)
		return
	}
	fmt.Println(string(body))
}
