package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func FileInfo(path string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Файл %s не существует\n", path)
			return
		}
		log.Fatal(err)
	}
	fmt.Println("Имя:", fileInfo.Name())
	fmt.Println("Размер:", fileInfo.Size())
	fmt.Println("Является директорией:", fileInfo.IsDir())
	fmt.Println("Время:", fileInfo.ModTime().Format("02.01.06 15:04:05"))
}

func main() {
	// подставим несуществующий файл
	FileInfo("file.go")
	// проверим файл main.go
	FileInfo("main.go")
}
