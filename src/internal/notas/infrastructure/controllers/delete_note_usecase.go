package controllers

import (
	"ARQ.HEX/src/internal/notas/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteNoteController struct {
	useCase *application.DeleteNote
}

func NewDeleteNoteController(useCase *application.DeleteNote) *DeleteNoteController{
	return &DeleteNoteController{useCase: useCase}
}

func (c *DeleteNoteController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.useCase.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Nota eliminada con éxito"})
}
