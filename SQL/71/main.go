package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM clients WHERE id = :id", sql.Named("id", 3))
	if err != nil {
		fmt.Println(err)
		return
	}
}
