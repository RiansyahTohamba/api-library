package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("file:librarycli.db"))
	if err != nil {
		log.Println(err)
	}
	// initDB(db)
	// get books
	printAll(db)
	// ?
	updateExample(db)
	printAll(db)

}

func updateExample(db *gorm.DB) {
	var book Book

	bookId1 := 100
	err := db.Where("id = ?", bookId1).First(&book).Error

	if err != nil {
		log.Println(err)
		return
	}
	book.Title = "Naruto 20"
	book.Price = 20000

	err = db.Save(&book).Error

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("book has been updated")
	}
}
func printAll(db *gorm.DB) {
	var books []Book
	err := db.Find(&books).Error
	if err != nil {
		log.Println(err)
	}

	fmt.Println(books)
}

func initDB(db *gorm.DB) {
	// db create schema?
	createSchema(db)
	// init books
	seedBook(db)
}

func createSchema(db *gorm.DB) {
	err := db.Exec(`CREATE TABLE IF NOT EXISTS books(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(255) NOT NULL,
		price REAL NOT NULL
	);

	CREATE TABLE IF NOT EXISTS authors(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(255) NOT NULL
	); `).Error

	if err != nil {
		log.Println(err)
	}
}
func seedBook(db *gorm.DB) {
	var books []Book

	books = append(books, Book{Title: "resep kentang goreng", Price: 100000.0})
	books = append(books, Book{Title: "resep mie goreng", Price: 200000.0})
	books = append(books, Book{Title: "resep nasi goreng", Price: 400000.0})

	for _, bk := range books {
		err := db.Create(&bk).Error
		if err != nil {
			log.Println(err)
		}
	}
	fmt.Println("book has inserted")
}

type Book struct {
	ID    int
	Title string
	Price float32
}
