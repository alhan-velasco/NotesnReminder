package routes

import (
	"github.com/gin-gonic/gin"
	"ARQ.HEX/src/internal/notas/infrastructure/controllers"
)

type NoteRoutes struct {
	CreateNoteController  *controllers.CreateNoteController
	GetAllNotesController *controllers.GetNotesController
	GetNoteByIDController *controllers.GetNoteByIDController
	UpdateNoteController  *controllers.UpdateNoteController
	DeleteNoteController  *controllers.DeleteNoteController
}

func NewNoteRoutes(
	createNoteController *controllers.CreateNoteController,
	getAllNotesController *controllers.GetNotesController,
	getNoteByIDController *controllers.GetNoteByIDController,
	updateNoteController *controllers.UpdateNoteController,
	deleteNoteController *controllers.DeleteNoteController,
) *NoteRoutes {
	return &NoteRoutes{
		CreateNoteController:  createNoteController,
		GetAllNotesController: getAllNotesController,
		GetNoteByIDController: getNoteByIDController,
		UpdateNoteController:  updateNoteController,
		DeleteNoteController:  deleteNoteController,
	}
}

func (r *NoteRoutes) AttachRoutes(router *gin.Engine) {
	notesGroup := router.Group("/notes")
	{
		notesGroup.POST("", r.CreateNoteController.Create)
		notesGroup.GET("", r.GetAllNotesController.GetAll)
		notesGroup.GET("/:id", r.GetNoteByIDController.GetByID)
		notesGroup.PUT("/:id", r.UpdateNoteController.Update)
		notesGroup.DELETE("/:id", r.DeleteNoteController.Delete)
	}
}
