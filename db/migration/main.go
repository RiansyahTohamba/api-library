package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db/library.db")
	if err != nil {
		log.Println(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS books(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			title VARCHAR(255) NOT NULL,
			price REAL NOT NULL
		);

		CREATE TABLE IF NOT EXISTS authors(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(255) NOT NULL
		);
		
		CREATE TABLE IF NOT EXISTS journals(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(255) NOT NULL
		);
	`)

	if err != nil {
		log.Println(err)
	}

	rows, err := db.Query("select * from books,authors,journals")

	if err != nil {
		log.Println(err)
	}
	fmt.Println(rows)
}
