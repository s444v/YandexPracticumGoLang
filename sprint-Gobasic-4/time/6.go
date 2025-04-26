package main

import (
	"fmt"
	"time"
)

// формат дней рождений
const layout = "02.01.2006"

func Age(birthday string) (int, error) {
	t, err := time.Parse(layout, birthday)
	var result int = 0
	timeNow := time.Now()
	if timeNow.Month() < t.Month() || (timeNow.Month() == t.Month() && timeNow.Day() > t.Day()) {
		result = -1
	}
	return int(time.Now().Year()) - int(t.Year()) + result, err
}

func main() {
	for _, v := range []string{"04.01.1969", "28.07.1995",
		"16.12.2007", "07.03.1947"} {
		age, err := Age(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(v, ":", age)
	}
}
