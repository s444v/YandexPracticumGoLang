package list

import (
	"testing"
)

func TestMin(t *testing.T) {
	// варианты последовательностей
	lists := [][]int{
		{5, -50, 67},
		{-7, 1, 43, 100},
		{3, 4},
		{10},
		{},
	}
	// ожидаемые значения для каждой последовательности
	mins := []int{-50, -7, 3, 10, 0}

	for i, list := range lists {
		if Min(list) != mins[i] {
			t.Error(i, ":", Min(list), "!=", mins[i])
		}
	}
}

func TestMul(t *testing.T) {
	listA := [][]int{
		{},
		{3},
		{-5, 6, 10},
	}
	listB := [][]int{
		{},
		{7},
		{21, 0, 4},
	}
	// ожидаемые слайсы results[i] = listA[i]*listB[i]
	results := [][]int{
		{},
		{21},
		{-105, 0, 40},
	}
	for i, list := range listA {
		c, err := Mul(list, listB[i])
		require.NoError(t, err)            // проверяем, что нет ошибки
		require.Len(t, c, len(results[i])) // проверяем длину резльтата
		// проверка значений
		for j, v := range c {
			require.Equal(t, results[i][j], v)
		}
	}
	// проверяем на возврат ошибки
	_, err := Mul([]int{22}, []int{33, 71})
	require.Error(t, err)
}
