package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleTime(res http.ResponseWriter, req *http.Request) {
	s := time.Now().Format("02.01.2006 15:04:05")
	res.Write([]byte(s))
}

func handleDate(res http.ResponseWriter, req *http.Request) {
	s := time.Now().Format("02.01.06")
	res.Write([]byte(s))
}

func handleMain(res http.ResponseWriter, req *http.Request) {
	s := fmt.Sprintf("Method: %s\nHost: %s\nPath: %s",
		req.Method, req.Host, req.URL.Path)
	res.Write([]byte(s))
}

func main() {
	http.HandleFunc("/time", handleTime)
	http.HandleFunc("/date", handleDate)
	http.HandleFunc("/", handleMain)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

// check
