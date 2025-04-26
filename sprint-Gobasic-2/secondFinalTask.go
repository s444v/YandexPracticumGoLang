package main

import (
	"fmt"
	"slices"
)

// Mode возвращает моды числовой последовательности.
// Напишите код функции
func Mode(a []int) ([]int, int) {
	var count int = 1
	var flag bool = false
	b := make([]int, 0)
	if len(a) == 0 {
		return b, count
	}
	for i, v := range a {
		tmp_count := 1
		found := slices.Contains(b, v)
		if found {
			continue
		}
		for j, k := range a {
			if i != j && v == k {
				flag = true
				tmp_count++
			}
		}
		if flag {
			flag = false
			b = append(b, v)
		}
		if tmp_count > count {
			count = tmp_count
		}
	}
	return b, count
}

func main() {
	lists := [][]int{
		{},
		{57},
		{78, -7},
		{99, 200, 0},
		{4, 4, 4, 3},
		{102, -7, 44, -7, 102},
		{82, -23, 1, 5, 98, 100},
		{100000, 90000, 20000, 20000, 20000, 22000, 25500, 22000},
	}
	modes := [][]int{
		{}, {}, {}, {},
		{4},
		{102, -7}, {},
		{20000, 22000},
	}
	mcount := []int{
		1, 1, 1, 1, 3, 2, 1, 3,
	}
	for i, list := range lists {
		mode, count := Mode(list)
		if len(mode) != len(modes[i]) {
			fmt.Printf("len mode %d: %v != %v'\n", i, modes[i], mode)
		} else {
			for j, v := range mode {
				if v != modes[i][j] {
					fmt.Printf("mode %d: %v != %v\n", i, modes[i], mode)
				}
			}
		}
		if count != mcount[i] {
			fmt.Printf("mcount %d: %d != %d\n", i, mcount[i], count)
		}
	}
	fmt.Println("Тестирование завершено")
}
