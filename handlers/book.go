package handlers

import (
	"api-library/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// create vs update
// update butuh ID, create tidak butuh id
type CreateBookRequest struct {
	Title string  `json:"title"`
	Price float32 `json:"price"`
}

type UpdatePostRequest struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float32 `json:"price"`
}

type bookAPI struct {
	// handler ini punya relasi 1-1 dengan bookRepository
	bookRepo *repository.BookRepository
}

func NewBook(bookRepo *repository.BookRepository) *bookAPI {
	return &bookAPI{bookRepo}
}

func (bk *bookAPI) Create(ctx *gin.Context) {
	var req CreateBookRequest = CreateBookRequest{}

	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail for binding request",
		})
	}

	err := bk.bookRepo.Create(req.Title, req.Price)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Data Book entered successfully",
		})

	} else {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "insert to db has been failed",
		})
	}
}

func (bk *bookAPI) List(ctx *gin.Context) {
	books, err := bk.bookRepo.FindAll()
	if err == nil {
		ctx.JSON(http.StatusOK, books)
	} else {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot retrieve books data",
		})
	}
}

func (bk *bookAPI) Show(ctx *gin.Context) {

	paramId := ctx.Param("id")
	bookId, err := strconv.Atoi(paramId)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "id must inserted through url paramater /:id",
		})
	}

	book, err := bk.bookRepo.FindByID(bookId)

	if err == nil {
		ctx.JSON(http.StatusOK, book)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot retrieve books data",
		})
	}
}

func (bk *bookAPI) Delete(ctx *gin.Context) {
	paramId := ctx.Param("id")
	bookId, err := strconv.Atoi(paramId)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "id must inserted through url paramater /:id",
		})
	}

	err = bk.bookRepo.Delete(bookId)

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "the book has been deleted",
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot delete books",
		})
	}
}

func (bk *bookAPI) Update(ctx *gin.Context) {
	var req = UpdatePostRequest{}

	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail for binding request",
		})
	}

	err := bk.bookRepo.Update(repository.Book{
		ID:    int64(req.ID),
		Title: req.Title,
		Price: req.Price,
	})

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "book has been updated",
		})
	} else {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot update value, check log",
		})
	}
}
