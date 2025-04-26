package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// слайс для хранения имен пользователей
var users []string

// handleUser обрабатывает запросы GET /user/{id}
func handleUser(w http.ResponseWriter, r *http.Request) {

	index, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || index < 0 || index >= len(users) {
		http.Error(w, "Некорректный идентификатор", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(users[index]))
}

// handleUsers обрабатывает GET /users
func handleUsers(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Ошибка при формировании JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

// handleAddUser обрабатывает POST /user
func handleAddUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Имя не определено", http.StatusBadRequest)
		return
	}
	users = append(users, name)

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strconv.Itoa(len(users) - 1)))
}

func main() {
	r := chi.NewRouter()
	r.Get("/user/{id}", handleUser)
	r.Get("/users", handleUsers)
	r.Post("/user", handleAddUser)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
