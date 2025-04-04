package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

type User struct {
	ID int `json:"id" yaml:"id"`
	// добавьте поля Name и Email
	Name  string `json:"name" yaml:"name"`
	Email string `json:"email" yaml:"email"`
}

func main() {
	// исходные данные в формате YAML
	yamlData := `
id: 2
name: Гофер
email: gopher@gophermate.com
`
	// промежуточная переменная типа User
	var user User
	// переменная для конвертации в JSON формат
	var jsonData []byte

	// добавьте десериализацию yamlData в user
	yaml.Unmarshal([]byte(yamlData), &user)
	// добавьте сериализацию user в jsonData
	jsonData, _ = json.MarshalIndent(user, "", "    ")
	fmt.Println(string(jsonData))
}
