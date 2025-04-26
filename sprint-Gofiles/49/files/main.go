package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// получаем текущую директорию
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Текущая директория:", curDir)
	// меняем текущую директорию на родительскую
	err = os.Chdir("..")
	// получаем новую текущую директорию
	curDir, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Новая текущая директория:", curDir)
}
