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

	productID := 1
	price := 700
	// обновление цены у продукта с заданным идентификатором
	// цена и идентификатор передаются через параметры запроса
	_, err = db.Exec("UPDATE products SET price = :price WHERE id = :id",
		sql.Named("price", price),
		sql.Named("id", productID))
	if err != nil {
		fmt.Println(err)
		return
	}
}
