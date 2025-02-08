package controllers

import (
	"ARQ.HEX/src/internal/recordatorio/application"
	"ARQ.HEX/src/internal/recordatorio/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type UpdateReminderController struct {
	useCase *application.UpdateReminder
}

func NewUpdateReminderController(useCase *application.UpdateReminder) *UpdateReminderController {
	return &UpdateReminderController{useCase: useCase}
}

func (c *UpdateReminderController) Update(ctx *gin.Context) {
	idstr := ctx.Param("id")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var request struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		DateTime    time.Time `json:"dateTime"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reminder := entities.Reminder{
		ID: int32(id),
		Title: request.Title,
		Description: request.Description,
		DateTime: request.DateTime,
	}

	if err := c.useCase.Execute(reminder); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Recordatorio actualizado con éxito"})
}
