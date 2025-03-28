package routes

import (
	"ARQ.HEX/src/internal/notas/application"
	"ARQ.HEX/src/internal/notas/application/services"
	"ARQ.HEX/src/internal/notas/infrastructure/controllers"
	"ARQ.HEX/src/internal/notas/infrastructure/handlers"
	"ARQ.HEX/src/internal/notas/infrastructure/dependencies/repository"
	"github.com/gin-gonic/gin"
	"database/sql"
)

type NotasDependencies struct {
	DB *sql.DB
}

func NewNotasDependencies(db *sql.DB) *NotasDependencies {
	return &NotasDependencies{DB: db}
}

func RegisterNoteRoutes(router *gin.Engine, d *NotasDependencies) {
	notaRepo := repository.NewNoteRepositoryImpl(d.DB)  

	noteService := services.NewNoteService(notaRepo)

	createNoteUseCase := application.NewCreateNoteUseCase(notaRepo)
	getAllNotesUseCase := application.NewGetNotesUseCase(notaRepo)
	getNoteByIDUseCase := application.NewGetNoteByIDUseCase(notaRepo)
	updateNoteUseCase := application.NewUpdateNoteUseCase(notaRepo)
	deleteNoteUseCase := application.NewDeleteNoteUseCase(notaRepo)

	createNoteController := controllers.NewCreateNoteController(createNoteUseCase)
	getAllNotesController := controllers.NewGetNotesController(getAllNotesUseCase)
	getNoteByIDController := controllers.NewGetNoteByIDController(getNoteByIDUseCase)
	updateNoteController := controllers.NewUpdateNoteController(updateNoteUseCase)
	deleteNoteController := controllers.NewDeleteNoteController(deleteNoteUseCase)

	notesGroup := router.Group("/notes")
	{
		notesGroup.POST("", createNoteController.Create)
		notesGroup.GET("", getAllNotesController.GetAll)
		notesGroup.GET("/:id", getNoteByIDController.GetByID)
		notesGroup.PUT("/:id", updateNoteController.Update)
		notesGroup.DELETE("/:id", deleteNoteController.Delete)

		notesGroup.GET("/new", handlers.GetNewNote(*noteService))
		notesGroup.GET("/deleted", handlers.GetDeletedNotes(*noteService))
		notesGroup.GET("/for-deletion", handlers.WaitForNoteDeletion(*noteService))
		notesGroup.GET("/wait-for-new", handlers.WaitForNewNote(*noteService))
	}
}
