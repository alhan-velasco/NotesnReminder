package controllers

import (
	"ARQ.HEX/src/internal/recordatorio/application"
	"ARQ.HEX/src/internal/recordatorio/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateReminderController struct {
	useCase *application.CreateReminder
}

func NewCreateReminderController(useCase *application.CreateReminder) *CreateReminderController {
	return &CreateReminderController{useCase: useCase}
}

func (c *CreateReminderController) Create(ctx *gin.Context) {
	var request struct {
		Title string `json:"title"`
		Description string `json:"description"`
		DateTime time.Time `json:"dateTime"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reminder := entities.Reminder{
		Title: request.Title,
		Description: request.Description,
		DateTime: request.DateTime,
	}

	if err := c.useCase.Execute(reminder); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Recordatorio creado con éxito"})
}
