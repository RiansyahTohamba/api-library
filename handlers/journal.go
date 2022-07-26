package handlers

import (
	"api-library/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type journalCreateRequest struct {
	Title string `json:"title"`
}

type journalUpdateRequest struct {
	Title string `json:"title"`
}
type journalAPI struct {
	jrepo *repository.JournalRepo
}

func NewJournal(jrepo *repository.JournalRepo) *journalAPI {
	return &journalAPI{jrepo}
}

func (ja *journalAPI) Create(ctx *gin.Context) {
	var req = journalCreateRequest{}
	if err := ctx.ShouldBind(&req); err != nil {

	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"request": req,
	})
}

func (ja *journalAPI) Update(ctx *gin.Context) {
	var req = journalUpdateRequest{}
	if err := ctx.ShouldBind(&req); err != nil {

	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"request": req,
	})
}

func (ja *journalAPI) Delete(ctx *gin.Context) {

}

func (ja *journalAPI) Show(ctx *gin.Context) {

}
func (ja *journalAPI) List(ctx *gin.Context) {

}
