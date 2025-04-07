package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("console.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	stdout := os.Stdout
	os.Stdout = f
	fmt.Println("Это вывод в файл consloe.txt")
	os.Stdout = stdout
	fmt.Println("Это вывод в терминал")
}
