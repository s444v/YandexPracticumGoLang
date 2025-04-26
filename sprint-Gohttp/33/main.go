package main

import (
	"fmt"
	"net/http"
	"time"
)

func mainHandle(res http.ResponseWriter, req *http.Request) {
	var out string

	if req.URL.Path == `/time` || req.URL.Path == `/time/` {
		out = time.Now().Format("02.01.2006 15:04:05")
	} else {
		out = fmt.Sprintf("Host: %s\nPath: %s\nMethod: %s",
			req.Host, req.URL.Path, req.Method)
	}
	res.Write([]byte(out))
}

func main() {
	http.HandleFunc(`/`, mainHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
