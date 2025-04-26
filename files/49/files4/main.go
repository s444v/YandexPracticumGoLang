package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Dir(path, shift string, size *int64) error {
	// получаем список файлов и поддиректорий
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		fi, err := file.Info()
		if err != nil {
			return err
		}
		if file.IsDir() {
			fmt.Printf("%s%s\n", shift, file.Name())
			// заходим внутрь директории
			err = Dir(filepath.Join(path, file.Name()), shift+"  ", size)
			if err != nil {
				return err
			}
			continue
		}
		// выводим информацию о файле
		*size += fi.Size()
		fmt.Printf("%s%s %s %d\n", shift, file.Name(),
			fi.ModTime().Format("02.01.06 15:04:05"), fi.Size())
	}
	return nil
}

func main() {
	var size int64
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Директория:", curDir)
	if err = Dir(curDir, "", &size); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%d\n", size)
}
