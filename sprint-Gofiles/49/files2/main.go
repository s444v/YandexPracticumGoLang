package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	// формируем относительное имя нужного файла
	dbName := filepath.Join("user", "data.db")
	// получаем его абсолютный путь
	dbPath, err := filepath.Abs(dbName)
	if err != nil {
		log.Fatal(err)
	}
	// получаем директорию файла
	fmt.Println("Директория:", filepath.Dir(dbPath))
	// получаем имя файла без пути
	fmt.Println("Имя файла:", filepath.Base(dbPath))
	// получаем расширение файла
	fmt.Println("Расширение:", filepath.Ext(dbPath))
}
