package geometry

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Answers struct {
	a      int
	b      int
	answer string
}

func TestHypotenuse(t *testing.T) {
	a := []Answers{
		{0, 0, "0.00"},
		{0, 9, "9.00"},
		{3, 4, "5.00"},
		{10, 21, "23.26"},
		{56, 37, "67.12"},
		{102, 67, "122.04"},
		{34, 17, "38.01"},
		{3456, 1089, "3623.51"},
		{478, 201, "518.54"},
	}
	for _, v := range a {
		h := Hypotenuse(float64(v.a), float64(v.b))
		assert.Equal(t, v.answer, fmt.Sprintf("%.2f", h))
	}

}
