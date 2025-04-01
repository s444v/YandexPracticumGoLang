package main

import (
	"fmt"
	"net/http"
)

func mainHandle(res http.ResponseWriter, req *http.Request) {
	s := fmt.Sprintf("%s/%s", req.Host, req.URL.Path)
	res.Write([]byte(s))
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
