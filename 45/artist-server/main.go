package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Artist описывает музыкальную группу.
type Artist struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Genre string   `json:"genre"`
	Songs []string `json:"songs"`
}

// переменная artists содержит пока один музыкальный коллектив
var artists = map[string]Artist{
	"30 seconds to Mars": {
		ID:    1,
		Name:  "30 seconds to Mars",
		Genre: "rock",
		Songs: []string{
			`The Kill`,
			`A Beautiful Lie`,
			`Attack`,
			`Live Like A Dream`,
		},
	},
}

// JSONHandler принимает значение из параметра band, ищет по нему в мапе группу, конвертирует
// данные из переменной band в JSON и выводит их в браузере.
func JSONHandler(w http.ResponseWriter, r *http.Request) {
	var band string

	// берем название группы из параметра `band`
	band = r.URL.Query().Get("band")

	// Проверяем POST-запрос или нет
	if r.Method == http.MethodPost {
		var artist Artist
		var buf bytes.Buffer
		// читаем тело запроса
		_, err := buf.ReadFrom(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// десериализуем JSON в Artist
		if err = json.Unmarshal(buf.Bytes(), &artist); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		artists[band] = artist
	}

	// если не POST-запрос, то всё как в предыдущем примере
	resp, err := json.Marshal(artists[band])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func main() {
	http.HandleFunc(`/`, JSONHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("ошибка запуска сервера: %s\n", err.Error())
		return
	}
}
