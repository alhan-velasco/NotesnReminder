package controllers

import (
	"ARQ.HEX/src/internal/recordatorio/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetReminderByIDController struct {
	useCase *application.GetReminderByID
}

func NewGetReminderByIDController(useCase *application.GetReminderByID) *GetReminderByIDController {
	return &GetReminderByIDController{useCase: useCase}
}

func (c *GetReminderByIDController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	reminder, err := c.useCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reminder)
}
