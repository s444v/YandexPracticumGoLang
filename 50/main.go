package main

import (
	"log"
	"os"
	"time"
)

func main() {
	fTime, err := os.OpenFile("run.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer fTime.Close()
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	fTime.Write([]byte(timeNow + "\n"))
}
