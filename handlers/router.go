package handlers

import (
	"api-library/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start(bookRepo *repository.BookRepository, journalRepo *repository.JournalRepo) {
	book := NewBook(bookRepo)
	journal := NewJournal(journalRepo)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))
	{
		router.POST("/book/", book.Create)
		router.GET("/books", book.List)
		router.GET("/book/:id", book.Show)
		router.PUT("/book/", book.Update)
		router.DELETE("/book/:id", book.Delete)
	}
	{
		router.POST("/journal/", journal.Create)
		router.GET("/journals", journal.List)
		router.GET("/journal/:id", journal.Show)
		router.PUT("/journal/", journal.Update)
		router.DELETE("/journal/:id", journal.Delete)
	}

	router.Run()

}
