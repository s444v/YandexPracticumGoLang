package main

import (
	"io"
	"net/http"
	"strconv"
	"strings"
)

var cafeList = map[string][]string{
	"moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент", "Ложка и вилка"},
	"tula":   []string{"Пир и мир", "Красиво есть не запретишь", "Поздний завтрак"},
}

func mainHandle(w http.ResponseWriter, req *http.Request) {
	var err error

	// если count не указан, то возвращается 25 записей
	count := 25
	countStr := req.FormValue("count")
	if countStr != "" {
		count, err = strconv.Atoi(countStr)
		if err != nil {
			http.Error(w, "incorrect count", http.StatusBadRequest)
			return
		}
	}
	city := req.FormValue("city")
	cafe, ok := cafeList[city]
	if !ok {
		http.Error(w, "unknown city", http.StatusBadRequest)
		return
	}
	if search := req.FormValue("search"); search != "" {
		var found []string

		for _, v := range cafe {
			if strings.Contains(strings.ToLower(v), strings.ToLower(search)) {
				found = append(found, v)
			}
		}
		cafe = found
	}
	count = min(count, len(cafe))
	answer := strings.Join(cafe[:count], ",")
	io.WriteString(w, answer)
}

func main() {
	http.HandleFunc(`/cafe`, mainHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
