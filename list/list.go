// list.go
package list

import "fmt"

func Min(list []int) int {
	var min int

	for _, v := range list {
		if v < min {
			min = v
		}
	}
	return min
}

func Mul(a, b []int) ([]int, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("different length of slices")
	}
	result := make([]int, len(a))
	for i, v := range a {
		result[i] = v * b[i]
	}
	return result, nil
}
