package controllers

import (
	"ARQ.HEX/src/internal/notas/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetNoteByIDController struct {
	useCase *application.GetNoteByID
}

func NewGetNoteByIDController(useCase *application.GetNoteByID) *GetNoteByIDController {
	return &GetNoteByIDController{useCase: useCase}
}

func (c *GetNoteByIDController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	note, err := c.useCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, note)
}
