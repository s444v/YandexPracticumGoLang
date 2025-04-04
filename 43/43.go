package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Artist описывает музыкальную группу.
type Artist struct {
	ID    int      `json:"-"`
	Name  string   `json:"name"`
	Genre string   `json:"genre,omitempty"`
	Songs []string `json:"songs"`
	// ID    int
	// Name  string
	// Genre string
	// Songs []string
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

// JSONHandler принимает значение из параметра id, ищет по нему в мапе группу, конвертирует
// данные из переменной artists в JSON и выводит их в браузере.
func JSONHandler(w http.ResponseWriter, r *http.Request) {
	band := r.URL.Query().Get("id")
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
