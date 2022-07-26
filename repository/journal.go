package repository

import "gorm.io/gorm"

type Journal struct {
	ID int
}

type JournalRepo struct {
	db *gorm.DB
}

func NewJournalRepository(db *gorm.DB) *JournalRepo {
	return &JournalRepo{db}
}

func Create() {

}
