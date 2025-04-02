package main

import (
	"fmt"
	"io"
	"net/http"
)

func mainHandle(w http.ResponseWriter, req *http.Request) {
	var answer string

	name := req.URL.Query().Get("name")
	if len(name) == 0 {
		answer = "Укажите имя заголовка в параметре: http://localhost:8080/?name=User-Agent"
	} else if v := req.Header.Get(name); len(v) > 0 {
		answer = fmt.Sprintf("%s: %s", name, v)
	} else {
		answer = fmt.Sprintf("Заголовок %s не определён", name)
	}
	io.WriteString(w, answer)
}

func main() {
	http.HandleFunc("/", mainHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
