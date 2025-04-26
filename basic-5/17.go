package main

import "fmt"

type Character struct {
	Name   string // имя
	Health int    // здоровье
	Speed  int    // скорость
	Power  int    // сила
	Woman  bool   // true, если женский персонаж
}

func (character Character) String() string {
	s := fmt.Sprintf(`Имя: %s
Здоровье: %d
Скорость: %d
Сила: %d
`, character.Name, character.Health, character.Speed, character.Power)
	return s
}

type Magician struct {
	Character
	Magic int
}

func (magician Magician) String() string {
	s := fmt.Sprintf("%sМагия: %d\n", magician.Character.String(), magician.Magic)
	return s
}

func main() {
	merlin := Magician{
		Character: Character{
			Name:   "Merlin",
			Health: 100,
			Speed:  250,
			Power:  400,
		},
		Magic: 700,
	}
	fmt.Println(merlin)
}
