package geometry

import "math"

// Hypotenuse вычисляет длину гипотенузы по длине двух катетов
func Hypotenuse(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}
