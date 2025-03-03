package controllers

import (
	"ARQ.HEX/src/internal/notas/application"
	"ARQ.HEX/src/internal/notas/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateNoteController struct {
	useCase *application.UpdateNote
}

func NewUpdateNoteController(useCase *application.UpdateNote) *UpdateNoteController {
	return &UpdateNoteController{useCase: useCase}
}

func (c *UpdateNoteController) Update(ctx *gin.Context) {
	idstr :=ctx.Param("id")

	id, err := strconv.ParseInt(idstr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var request struct {
		Title string `json:"title"`
		Content string `json:"content"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note := entities.Note{
		ID: int(id),
		Title: request.Title,
		Content: request.Content,
	}

	if err := c.useCase.Execute(note); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Nota actualizada con éxito"})
}
