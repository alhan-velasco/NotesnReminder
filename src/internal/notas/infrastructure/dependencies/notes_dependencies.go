package dependencies

import (
	"ARQ.HEX/src/internal/notas/application"
	"ARQ.HEX/src/internal/notas/infrastructure/controllers"
	"ARQ.HEX/src/internal/notas/infrastructure/dependencies/repository"
	"ARQ.HEX/src/internal/notas/infrastructure/dependencies/routes"
	"database/sql"
)

type NotasDependencies struct {
	DB *sql.DB
}

func NewNotasDependencies(db *sql.DB) *NotasDependencies {
	return &NotasDependencies{DB: db}
}

func (d *NotasDependencies) GetRoutes() *routes.NoteRoutes {
	notaRepo := repository.NewNoteRepositoryImpl(d.DB)

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

	return routes.NewNoteRoutes(
		createNoteController,
		getAllNotesController,
		getNoteByIDController,
		updateNoteController,
		deleteNoteController,
	)
}