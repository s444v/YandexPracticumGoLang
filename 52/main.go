package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(fileName string, process func(int, string)) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		process(lineNum, scanner.Text())
		lineNum++
	}
	return scanner.Err()
}
func main() {
	process := func(num int, line string) {
		fmt.Printf("%02d: %s\n", num, line)
	}
	ReadLines("main.go", process)
}
