package main

import (
	"fmt"
	"net/url"
)

func FormValues(achievement string) url.Values {
	// переделайте код функции с использованием переменной типа url.Values
	// и методов Set() и Add()
	values := url.Values{}
	values.Set("name", "Вася")
	values.Set("nick", "superstar")
	values.Set("achieves", "cool")
	values.Add("achieves", "best")
	values.Add("achieves", achievement)
	return values
}

func main() {
	vals := FormValues("80 level")

	fmt.Println(vals["name"])
	fmt.Println(vals["nick"])
	fmt.Println(vals["achieves"])
}
