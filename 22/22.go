package main

import (
	"fmt"
)

// Hero описывает героя с общими полями для всех классов
// в этой структуре ничего не поменялось
type Hero struct {
	Health int // здоровье
	Damage int // урон
	Def    int // защита
}

// Attack возвращает строку с информацией о нанесённом уроне
func (h Hero) Attack() string {
	return fmt.Sprint("Ваш персонаж нанёс урон, равный ", h.Damage)
}

// Defense возвращает строку с информацией о заблокированном уроне
func (h Hero) Defense() string {
	return fmt.Sprint("Ваш персонаж заблокировал ", h.Def, " урона")
}

// в следующие три структуры добавим урон Damage

// Warrior описывает класс «Воин»
type Warrior struct {
	Hero             // встроенная структура Hero
	Name      string // имя
	ClassName string // имя класса
	Damage    int    // урон для класса «Воин»
	// добавьте поле ниже
	Def int
}

// переопределяем метод Attack() для воина
func (w Warrior) Attack() string {
	return fmt.Sprintf("%s класса %s нанёс %d урона", w.Name, w.ClassName, w.Damage+w.Hero.Damage)
}

// переопределите метод Defense() для воина здесь
func (w Warrior) Defense() string {
	return fmt.Sprintf("%s класса %s заблокировал %d урона", w.Name, w.ClassName, w.Def+w.Hero.Def)
}

// Mage описывает класс «Маг»
type Mage struct {
	Hero             // встроенная структура Hero
	Name      string // имя
	ClassName string // имя класса
	Damage    int    // урон для класса «Маг»
	// добавьте поле ниже
	Def int
}

// переопределяем метод Attack() для мага
func (m Mage) Attack() string {
	return fmt.Sprintf("%s класса %s нанёс %d урона", m.Name, m.ClassName, m.Damage+m.Hero.Damage)
}

// переопределите метод Defense() для мага здесь
func (m Mage) Defense() string {
	return fmt.Sprintf("%s класса %s заблокировал %d урона", m.Name, m.ClassName, m.Def+m.Hero.Def)
}

// Healer описывает класс «Лекарь»
type Healer struct {
	Hero             // встроенная структура Hero
	Name      string // имя
	ClassName string // имя класса
	Damage    int    // урон для класса «Лекарь»
	// добавьте поле ниже
	Def int
}

// переопределяем метод Attack() для лекаря
func (h Healer) Attack() string {
	return fmt.Sprintf("%s класса %s нанёс %d урона", h.Name, h.ClassName, h.Damage+h.Hero.Damage)
}

// переопределите метод Defense() для лекаря здесь
func (h Healer) Defense() string {
	return fmt.Sprintf("%s класса %s заблокировал %d урона", h.Name, h.ClassName, h.Def+h.Hero.Def)
}

func main() {
	// общие параметры для всех классов в структуре Hero
	hero := Hero{Health: 100, Damage: 30, Def: 20}

	// воин
	warrior := Warrior{Hero: hero, Name: "Арагорн", ClassName: "Воин", Damage: 20, Def: 30}
	// воин атакует
	fmt.Println(warrior.Attack())
	// воин защищается
	fmt.Println(warrior.Defense())

	// маг
	mage := Mage{Hero: hero, Name: "Мерлин", ClassName: "Маг", Damage: 30, Def: 10}
	// маг атакует
	fmt.Println(mage.Attack())
	// маг защищается
	fmt.Println(mage.Defense())

	// лекарь
	healer := Healer{Hero: hero, Name: "Елена Малышева", ClassName: "Лекарь", Damage: 10, Def: 20}
	// лекарь атакует
	fmt.Println(healer.Attack())
	// лекарь защищается
	fmt.Println(healer.Defense())
}
