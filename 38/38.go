package main

import (
	"fmt"
	"net/http"
)

const pattern = `<!DOCTYPE html>
  <html lang="ru"><head>
  <meta charset="utf-8" />
  <title>Тестовый сервер</title>
  </head>
<body>%s</body></html>`

func mainHandle(w http.ResponseWriter, req *http.Request) {
	// устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, pattern,
			fmt.Sprintf("Сервер не поддерживает %s запросы", req.Method))
		return
	}
	fmt.Fprintf(w, pattern, "Получен GET-запрос")
}

func main() {
	http.HandleFunc("/", mainHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
