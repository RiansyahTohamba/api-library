package handlers

import (
	"api-library/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type journalAPI struct {
	jrepo *repository.JournalRepo
}

func NewJournal(jrepo *repository.JournalRepo) *journalAPI {
	return &journalAPI{jrepo}
}

func (ja *journalAPI) Create(ctx *gin.Context){

}

func (ja *journalAPI) Update(ctx *gin.Context){

}

func (ja *journalAPI) Delete(ctx *gin.Context)  {
	
}


func Start(bookRepo *repository.BookRepository, journalRepo *repository.JournalRepo) {
	book := NewBook(bookRepo)
	journal := NewJournal(journalRepo)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	// create book
	router.POST("/book/", book.Create)
	// get list of books
	router.GET("/book/list", book.List)
	// get one book
	router.GET("/book/:id", book.Show)
	// update book
	router.PUT("/book/", book.Update)
	// delete book
	router.DELETE("/book/:id", book.Delete)
	router.GET("",journal.)
	router.Run()

}
