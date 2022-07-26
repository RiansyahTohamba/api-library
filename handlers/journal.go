package handlers

import (
	"api-library/repository"

	"github.com/gin-gonic/gin"
)

type journalAPI struct {
	jrepo *repository.JournalRepo
}

func NewJournal(jrepo *repository.JournalRepo) *journalAPI {
	return &journalAPI{jrepo}
}

func (ja *journalAPI) Create(ctx *gin.Context) {

}

func (ja *journalAPI) Update(ctx *gin.Context) {

}

func (ja *journalAPI) Delete(ctx *gin.Context) {

}

func (ja *journalAPI) Show(ctx *gin.Context) {

}
func (ja *journalAPI) List(ctx *gin.Context) {

}
