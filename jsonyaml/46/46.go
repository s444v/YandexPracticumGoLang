package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"-"`
	active   bool
}

func main() {
	user := User{
		ID:       1,
		Name:     "Гофер",
		Email:    "gopher@gophermate.com",
		Password: "I4mG0ph3R",
		active:   true,
	}

	// преобразуйте user в JSON формат
	resp, err := json.Marshal(user)
	if err != nil {
		fmt.Print(err)
		return
	}
	var newUser User
	if err = json.Unmarshal(resp, &newUser); err != nil {
		fmt.Print(err)
		return
	}

	// десериализуйте данные из JSON формата в переменную newUser

	fmt.Println(newUser)
}
