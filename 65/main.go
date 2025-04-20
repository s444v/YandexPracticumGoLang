package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type Product struct {
	ID      int
	Product string
	Price   int
}

func (p Product) String() string {
	return fmt.Sprintf("ID: %d, Product: %s, Price: %d", p.ID, p.Product, p.Price)
}

func main() {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, product, price FROM products")
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		product := Product{}
		err := rows.Scan(&product.ID, &product.Product, &product.Price)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(product)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return
	}
}
