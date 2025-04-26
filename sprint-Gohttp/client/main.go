package main

import (
	"encoding/json"
	"net/http"
)

type Artist struct {
	ID    string   `json:"id"`    // id коллектива
	Name  string   `json:"name"`  // название группы
	Born  string   `json:"born"`  // год основания группы
	Genre string   `json:"genre"` // жанр
	Songs []string `json:"songs"` // популярные песни, это слайс строк,
	// так как песен может быть несколько
}

// ...
var artists = map[string]Artist{
	"1": {
		ID:    "1",
		Name:  "30 Seconds To Mars",
		Born:  "1998",
		Genre: "alternative",
		Songs: []string{
			"The Kill",
			"A Beautiful Lie",
			"Attack",
			"Live Like A Dream",
		},
	},
	"2": {
		ID:    "2",
		Name:  "Garbage",
		Born:  "1994",
		Genre: "alternative",
		Songs: []string{
			"Queer",
			"Shut Your Mouth",
			"Cup of Coffee",
			"Til the Day I Die",
		},
	},
}

func getArtists(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(artists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
