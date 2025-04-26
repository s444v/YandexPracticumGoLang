package main

import (
	"fmt"
)

// Hero описывает героя с общими полями для всех классов
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

// Inventory описывает инвентарь
// Добавьте структуру Inventory ниже
type Inventory struct {
	Items map[string]int
}

// Take добавляет предмет в инвентарь
// Добавьте метод Take() для Inventory тут
/*
Метод Take() будет добавлять в мапу предмет. Нужно будет проверить, есть ли уже такой предмет в мапе.
Если есть — увеличить количество на 1 и вывести сообщение «Вы положили <предмет> в инвентарь. Количество: <количество_предметов>».
Если нет, то количество этого предмета будет 1 и нужно будет вывести сообщение «Вы положили <предмет> в инвентарь». Вспомните, как работать с мапами.
*/
func (i *Inventory) Take(item string) string {
	_, ok := i.Items[item]
	if ok {
		i.Items[item] += 1
		return fmt.Sprint("Вы положили ", item, " в инвентарь. Количество: ", i.Items[item], " .")
	} else {
		i.Items[item] = 1
		return fmt.Sprint("Вы положили ", item, " в инвентарь. ")
	}
}

// Drop удаляет предмет из инвентаря
// Добавьте метод Drop() для Inventory тут
/*
Метод Drop() будет удалять один предмет из инвентаря. Нужно проверить, есть ли вообще такой предмет в инвентаре.
Если нет, вывести сообщение «У вас нет предмета <предмет>». Также нужно будет проверить, сколько таких предметов.
Если всего один, нужно полностью удалить его из инвентаря и вывести сообщение «Вы выбросили <предмет>».
Если несколько, то нужно уменьшить количество этого предмета на 1 и вывести сообщение «Вы выбросили <предмет>», а в новой строке — «Осталось <количество_предметов>».
*/
func (i *Inventory) Drop(item string) string {
	count, ok := i.Items[item]
	if ok && count == 1 {
		delete(i.Items, item)
		return fmt.Sprint("Вы выбросили ", item, " .")
	} else if ok && count != 1 {
		i.Items[item] -= 1
		return fmt.Sprint("Вы выбросили ", item, " .\n", "Осталось ", i.Items[item])
	} else {
		return fmt.Sprint("У вас нет предмета ", item, " .")
	}
}

// Warrior описывает класс «Воин»
type Warrior struct {
	Hero
	Inventory
	Name      string // имя героя
	ClassName string // имя класса
}

// Buff — специальное умение для Warrior
func (w *Warrior) Buff() string {
	w.Def += 20
	return fmt.Sprintf("%s класса %s увеличил свою защиту.\nЗащита %s теперь %d.\n",
		w.Name,
		w.ClassName,
		w.Name,
		w.Def,
	)
}

// Mage описывает класс «Маг»
type Mage struct {
	Hero
	Inventory
	Name      string // имя героя
	ClassName string // имя класса
}

// Buff — специальное умение для Mage
func (m *Mage) Buff() string {
	m.Damage += 30
	return fmt.Sprintf("%s класса %s усилил свою атаку.\nАтака %s теперь %d.\n",
		m.Name,
		m.ClassName,
		m.Name,
		m.Damage,
	)
}

// Healer описывает класс «Лекарь»
type Healer struct {
	Hero
	Inventory
	Name      string // имя героя
	ClassName string // имя класса
}

// Buff — специальное умение для Healer
func (h *Healer) Buff() string {
	h.Health += 50
	return fmt.Sprintf("%s класса %s увеличил своё здоровье.\nЗдоровье %s теперь %d.\n",
		h.Name,
		h.ClassName,
		h.Name,
		h.Health,
	)
}

func main() {
	// Общая структура для всех классов с полями «Здоровье», «Урон» и «Защита»
	hero := Hero{Health: 100, Damage: 30, Def: 20}

	// Маг
	// Инвентарь для мага
	mageInventory := Inventory{make(map[string]int)}

	// Структура для мага
	mage := Mage{Hero: hero, Inventory: mageInventory, Name: "Мерлин", ClassName: "Маг"}

	// Представимся
	fmt.Println("Я", mage.Name, "класса", mage.ClassName)

	// Маг атакует
	fmt.Println(mage.Attack())
	// Маг защищается
	fmt.Println(mage.Defense())
	// Маг баффается
	fmt.Println(mage.Buff())
	// Маг кладёт посох в инвентарь
	fmt.Println(mage.Take("Посох"))
	// Посох не понравился, выбрасывает посох
	fmt.Println(mage.Drop("Посох"))

	// Разделим вывод для персонажей
	fmt.Println()

	// Воин

	// Инвентарь для воина
	warriorInventory := Inventory{make(map[string]int)}

	// Структура для воина
	warrior := Warrior{Hero: hero, Inventory: warriorInventory, Name: "Арагорн", ClassName: "Воин"}

	// Представимся
	fmt.Println("Я", warrior.Name, "класса", warrior.ClassName)

	// Воин атакует
	fmt.Println(warrior.Attack())
	// Воин защищается
	fmt.Println(warrior.Defense())
	// Воин баффается
	fmt.Println(warrior.Buff())
	// Воин кладёт в инвентарь меч
	fmt.Println(warrior.Take("Меч"))
	// Воин кладёт в инвентарь шлем
	fmt.Println(warrior.Take("Шлем"))
	// Воин кладёт в инвентарь наручи
	fmt.Println(warrior.Take("Наручи"))
	// Ещё один шлем в инвентарь
	fmt.Println(warrior.Take("Шлем"))
	// Много шлемов, один выкинул
	fmt.Println(warrior.Drop("Шлем"))
	// Попробовал выкинуть сапоги, но их же и так нет
	fmt.Println(warrior.Drop("Сапоги"))

	// Разделим вывод для персонажей
	fmt.Println()

	// Лекарь
	// инвентарь для лекаря
	healerInventory := Inventory{make(map[string]int)}

	// Структура для лекаря
	healer := Healer{Hero: hero, Inventory: healerInventory, Name: "Елена Малышева", ClassName: "Лекарь"}

	// Представимся
	fmt.Println("Я", healer.Name, "класса", healer.ClassName)

	// Лекарь атакует
	fmt.Println(healer.Attack())
	// Лекарь защищается
	fmt.Println(healer.Defense())
	// Лекарь накладывает на себя заклинание, или бафф
	fmt.Println(healer.Buff())
	// Лекарь кладёт в инвентарь амулет
	fmt.Println(healer.Take("Амулет"))
	// Лекарь кладёт в инвентарь плащ
	fmt.Println(healer.Take("Плащ"))
}
