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

	var (
		product string
		price   int
	)
	row := db.QueryRow("SELECT product, price FROM products WHERE id = :id", sql.Named("id", 1))
	err = row.Scan(&product, &price)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(product, price)
}
