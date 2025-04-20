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

	rows, err := db.Query("SELECT product FROM products WHERE price > :price", sql.Named("price", 500))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product string

		err := rows.Scan(&product)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(product)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return
	}
}
