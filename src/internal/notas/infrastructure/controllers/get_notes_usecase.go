package controllers

import (
	"ARQ.HEX/src/internal/notas/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetNotesController struct {
	useCase *application.GetNotes
}

func NewGetNotesController(useCase *application.GetNotes) *GetNotesController {
	return &GetNotesController{useCase: useCase}
}

func (c *GetNotesController) GetAll(ctx *gin.Context) {
	notes, err := c.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, notes)
}
