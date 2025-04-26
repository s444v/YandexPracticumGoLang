package main

import (
	"fmt"
)

// Person описывает персонажа
type Person struct {
	Name string // имя
	Age  int    // возраст
}

// сделайте из этой функции метод для Person
func (p Person) withName() string {
	return p.Name
}

// сделайте из этой функции метод для Person
func (p Person) withAge() int {
	return p.Age
}

func main() {
	p := Person{"Вася", 32}

	name := p.withName()

	fmt.Println("Имя персонажа", name)

	age := p.withAge()

	fmt.Println("Возраст персонажа:", age)
}
