package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Client struct {
	ID       int
	FIO      string
	Login    string
	Birthday string
	Email    string
}

func insertClient(db *sql.DB, client *Client) (int64, error) {
	res, err := db.Exec("INSERT INTO clients (fio, login, birthday, email) VALUES (:fio, :login, :birthday, :email)",
		sql.Named("fio", client.FIO),
		sql.Named("login", client.Login),
		sql.Named("birthday", client.Birthday),
		sql.Named("email", client.Email))
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func updateClientLogin(db *sql.DB, login string, id int64) error {
	_, err := db.Exec("UPDATE clients SET login = :login WHERE id = :id",
		sql.Named("login", login),
		sql.Named("id", id))
	return err
}

func deleteClient(db *sql.DB, id int64) error {
	_, err := db.Exec("DELETE FROM clients WHERE id = :id", sql.Named("id", id))
	return err
}

func selectClient(db *sql.DB, id int64) (Client, error) {
	client := Client{}
	row := db.QueryRow("SELECT * FROM clients WHERE id = :id", sql.Named("id", id))
	err := row.Scan(&client.ID, &client.FIO, &client.Login, &client.Birthday, &client.Email)
	if err != nil {
		return client, err
	}
	return client, err
}

// func main() {
// 	db, err := sql.Open("sqlite", "demo.db")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer db.Close()
// 	client := Client{
// 		FIO:      "Попов Степан Степанович",
// 		Login:    "loginglogin",
// 		Birthday: "20240101",
// 		Email:    "popov@mail.ru",
// 	}
// 	id, err := insertClient(db, &client)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	err = selectClient(db, id)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	err = updateClientLogin(db, "lulululu", id)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	err = selectClient(db, id)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	err = deleteClient(db, id)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	err = selectClient(db, id)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }
