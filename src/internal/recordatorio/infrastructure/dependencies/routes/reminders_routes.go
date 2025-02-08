package routes

import (
	"github.com/gin-gonic/gin"
	"ARQ.HEX/src/internal/recordatorio/infrastructure/controllers"
)

type ReminderRoutes struct {
	CreateReminderController  *controllers.CreateReminderController
	GetAllRemindersController *controllers.GetRemindersController
	GetReminderByIDController *controllers.GetReminderByIDController
	UpdateReminderController  *controllers.UpdateReminderController
	DeleteReminderController  *controllers.DeleteReminderController
}

func NewReminderRoutes(
	createReminderController *controllers.CreateReminderController,
	getAllRemindersController *controllers.GetRemindersController,
	getReminderByIDController *controllers.GetReminderByIDController,
	updateReminderController *controllers.UpdateReminderController,
	deleteReminderController *controllers.DeleteReminderController,
) *ReminderRoutes {
	return &ReminderRoutes{
		CreateReminderController:  createReminderController,
		GetAllRemindersController: getAllRemindersController,
		GetReminderByIDController: getReminderByIDController,
		UpdateReminderController:  updateReminderController,
		DeleteReminderController:  deleteReminderController,
	}
}

func (r *ReminderRoutes) AttachRoutes(router *gin.Engine) {
	remindersGroup := router.Group("/reminders")
	{
		remindersGroup.POST("", r.CreateReminderController.Create)
		remindersGroup.GET("", r.GetAllRemindersController.GetAll)
		remindersGroup.GET("/:id", r.GetReminderByIDController.GetByID)
		remindersGroup.PUT("/:id", r.UpdateReminderController.Update)
		remindersGroup.DELETE("/:id", r.DeleteReminderController.Delete)
	}
}
