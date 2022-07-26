package handlers

import "github.com/gin-gonic/gin"

type authorAPI struct {
}

func NewAuthor() *authorAPI {
	return &authorAPI{}
}
func (at *authorAPI) Show(ctx *gin.Context) {

}
