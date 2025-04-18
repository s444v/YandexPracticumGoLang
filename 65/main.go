package main

import (
	"database/sql"
	"fmt"

	// импортируйте нужный пакет
	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "school.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()
}
