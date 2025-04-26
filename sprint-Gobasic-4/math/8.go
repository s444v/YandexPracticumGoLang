package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	f := math.Sqrt(25)
	s := strings.Join([]string{"sqrt(25)", "равен"}, " ")
	fmt.Println(s, f)
}
