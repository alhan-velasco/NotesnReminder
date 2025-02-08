package controllers

import (
	"ARQ.HEX/src/internal/recordatorio/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteReminderController struct {
	useCase *application.DeleteReminder
}

func NewDeleteReminderController(useCase *application.DeleteReminder) *DeleteReminderController {
	return &DeleteReminderController{useCase: useCase}
}

func (c *DeleteReminderController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.useCase.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Recordatorio eliminado con éxito"})
}
