package main

import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("Запускаем сервер")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }
    fmt.Println("Завершаем работу")
}