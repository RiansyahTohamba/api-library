package main

import (
	"api-library/db"
	"api-library/handlers"
	"api-library/repository"
)

func main() {
	// db conn
	dbConn := db.GetConn()
	// repository
	bookRepo := repository.NewBookRepository(dbConn)
	jrepo := repository.NewJournalRepository(dbConn)
	// := repository.NewJournal
	handlers.Start(bookRepo, jrepo)
}
