package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func PrintAnswer(w *http.Response) {
	// читаем тело ответа
	body, err := io.ReadAll(w.Body)
	w.Body.Close()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println(string(body))
}

func main() {
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	PrintAnswer(resp)

	resp, err = http.PostForm("http://localhost:8080/", url.Values{
		"email": {"my@my-best-site.ru"},
		"name":  {"Василий"},
	})
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	PrintAnswer(resp)
}
