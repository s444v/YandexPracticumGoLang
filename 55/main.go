package main

import (
	"fmt"
	"math"
	"math/rand"
)

func Winner(rnd *rand.Rand, bets map[int]string) string {
	//
	generatedNumber := rnd.Int63() % 100
	diff, key := 100, 0
	for k, _ := range bets {
		if math.Abs(float64(int(generatedNumber)-(k))) < float64(diff) {
			diff = int(generatedNumber) - k
			key = k
		}
	}
	return bets[key]
}

func main() {
	rnd := rand.New(rand.NewSource(1001))
	bets := map[int]string{
		20: "Маша",
		34: "Игорь",
		77: "Олег",
		51: "Света",
		2:  "Саша",
	}
	fmt.Println("Победитель:", Winner(rnd, bets))
}
