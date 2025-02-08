package controllers

import (
	"ARQ.HEX/src/internal/notas/application"
	"ARQ.HEX/src/internal/notas/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateNoteController struct {
	useCase *application.CreateNoteUseCase
}

func NewCreateNoteController(useCase *application.CreateNoteUseCase) *CreateNoteController{
	return &CreateNoteController{useCase: useCase}
}

func (c *CreateNoteController) Create(ctx *gin.Context) {
	var request struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note := entities.Note{
		Title:  request.Title,
		Content: request.Content,
	}

	if err := c.useCase.Execute(note); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Nota creada con éxito"})
}
