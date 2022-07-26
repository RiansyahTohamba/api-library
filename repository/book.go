package repository

import (
	"errors"

	"gorm.io/gorm"
)

type Book struct {
	ID    int64
	Title string
	Price float32
}

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db}
}

// GET ALL
func (br *BookRepository) FindAll() ([]Book, error) {
	var books []Book
	err := br.db.Find(&books).Error
	return books, err
}

// SHOW
func (br *BookRepository) FindByID(bookId int) (Book, error) {
	var bk Book
	err := br.db.Where("id = ?", bookId).First(&bk).Error
	return bk, err
}

// CREATE
func (br *BookRepository) Create(title string, price float32) error {
	book := Book{Title: title, Price: price}
	err := br.db.Create(&book).Error
	return err
}

// UPDATE
func (br *BookRepository) Update(bookInput Book) error {
	var book Book
	err := br.db.Where("id = ?", bookInput.ID).First(&book).Error

	if err != nil {
		return errors.New("record not found")
	}

	book.Title = bookInput.Title
	book.Price = bookInput.Price

	err = br.db.Save(&book).Error
	return err
}

// DELETE
func (br *BookRepository) Delete(bookId int) error {
	err := br.db.Delete(&Book{}, bookId).Error
	return err
}
