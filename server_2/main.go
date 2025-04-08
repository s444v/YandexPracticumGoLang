package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var users []string

type User struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

func getUserName(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, err := strconv.Atoi(id)
	if i > len(users) || i < 0 || err != nil {
		http.Error(w, "Пользователь не найден", http.StatusBadRequest)
		return
	}
	resp, err := json.Marshal(User{i, users[i]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var userslol []User
	for i, v := range users {
		userslol = append(userslol, User{i, v})
	}
	resp, err := json.Marshal(userslol)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func postUser(w http.ResponseWriter, r *http.Request) {
	var user User
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.Unmarshal(buf.Bytes(), &user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users = append(users, user.Name)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func main() {
	r := chi.NewRouter()
	r.Get("/user/{id}", getUserName)
	r.Get("/users", getUsers)
	r.Post("/user", postUser)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
