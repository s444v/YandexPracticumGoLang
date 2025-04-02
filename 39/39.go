package main

import (
	"net/http"
)

const pattern = `<!DOCTYPE html>
  <html lang="ru"><head>
  <meta charset="utf-8" />
  <title>Тестовый сервер</title>
  </head>
<body>%s</body></html>`

func mainHandle(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Сервер не поддерживает "+req.Method, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte(req.FormValue("name")))
}

func main() {
	http.HandleFunc("/", mainHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
