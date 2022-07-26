package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetConn() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db/library.db"))
	if err != nil {
		log.Println(err)
	}
	return db
}
