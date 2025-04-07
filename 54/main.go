package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// LogEntry описывает структуру записи лог-файла
type LogEntry struct {
	// интересует только одно поле
	IP string `json:"ip"`
}

func main() {
	// открываем файл лога
	logFile, err := os.Open("access.log")
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// создаем мапу для подсчета запросов по IP
	counts := make(map[string]int)

	// создаем сканер для построчного чтения лог-файла
	scanner := bufio.NewScanner(logFile)

	// читаем файл построчно
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		// декодируем JSON-строку в структуру LogEntry
		var item LogEntry
		err := json.Unmarshal([]byte(scanner.Text()), &item)
		if err != nil {
			log.Printf("ошибка парсинга: %v. Строка: %s\n", err, line)
			continue
		}
		// увеличиваем счетчик для данного IP
		counts[item.IP]++
	}

	// проверяем ошибки во время чтения
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// создаём файл для записи результатов
	resFile, err := os.Create("result.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer resFile.Close()

	for ip, count := range counts {
		_, err := fmt.Fprintf(resFile, "%s %d\n", ip, count)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Сохранено IP-адресов:", len(counts))
}
