package controllers

import (
	"ARQ.HEX/src/internal/recordatorio/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetRemindersController struct {
	useCase *application.GetReminder
}

func NewGetRemindersController(useCase *application.GetReminder) *GetRemindersController {
	return &GetRemindersController{useCase: useCase}
}

func (c *GetRemindersController) GetAll(ctx *gin.Context) {
	reminders, err := c.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reminders)
}
