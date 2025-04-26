package main

import (
	"fmt"
)

// Hero описывает главного героя игры
type Hero struct {
	Name   string
	Health int
	Damage int
	Def    int
}

// ниже напишите метод Defense() для структуры Hero
func (h Hero) Defense() {
	fmt.Println(h.Name, "заблокировал", h.Def, "единиц урона")
}

// ниже напишите метод Special() для структуры Hero
func (h *Hero) Special(healthpoints int) {
	fmt.Println("Количество здоровья было", h.Health)
	h.Health += healthpoints
	fmt.Println("Количество здоровья стало", h.Health)
}

func main() {
	myHero := Hero{Name: "Артур", Health: 100, Damage: 30, Def: 20}
	myHero.Defense()
	myHero.Special(30)
}
