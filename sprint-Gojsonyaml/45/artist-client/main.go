package main

import (
	"bytes"
	"net/http"
)

func main() {
	http.Post(`http://localhost:8080/?band=Garbage`, `application/json`,
		// создаём слайс байтов, в который записываем JSON
		// мы специально указали ключи в разных регистрах, чтобы посмотреть, правильно ли сконвертируются данные
		bytes.NewBufferString(`{"ID": 10, "name": "Garbage", "Genre":"rock", "SongS":["Only Happy When It Rains", "Supid Girl", "Push It"]}`))
}
