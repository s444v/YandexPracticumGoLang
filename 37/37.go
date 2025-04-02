package main

import (
	"fmt"
	"io"
	"net/http"
)

func mainHandle(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		fmt.Fprintf(w, "Email: %s\nName: %s",
			req.PostFormValue("email"), req.PostFormValue("name"))
		return
	}
	io.WriteString(w, "Отправьте POST запрос с параметрами email и name")
}

func main() {
	http.HandleFunc("/", mainHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
