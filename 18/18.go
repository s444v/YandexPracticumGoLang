package main

import "fmt"

type Hero struct {
	Name     string
	Health   int
	Location string
}

// ниже объявите метод
func (h *Hero) MoveTo(newLocation string) {
	h.Location = newLocation
	fmt.Println("Артур переместился в " + h.Location)
}

func main() {
	myHero := Hero{Name: "Артур", Health: 100}
	// вызовите метод для myHero
	myHero.MoveTo("тронный зал")
}
