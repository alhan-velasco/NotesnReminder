package dependencies

import (
	"ARQ.HEX/src/internal/recordatorio/application"
	"ARQ.HEX/src/internal/recordatorio/infrastructure/controllers"
	"ARQ.HEX/src/internal/recordatorio/infrastructure/dependencies/repository"
	"ARQ.HEX/src/internal/recordatorio/infrastructure/dependencies/routes"
	"database/sql"
)

type RecordatoriosDependencies struct {
	DB *sql.DB
}

func NewRecordatoriosDependencies(db *sql.DB) *RecordatoriosDependencies {
	return &RecordatoriosDependencies{DB: db}
}

func (d *RecordatoriosDependencies) GetRoutes() *routes.ReminderRoutes {
	reminderRepo := repository.NewReminderRepositoryImpl(d.DB)

	createReminderUseCase := application.NewCreateReminder(reminderRepo)
	getAllRemindersUseCase := application.NewGetReminderUseCase(reminderRepo)
	getReminderByIDUseCase := application.NewGetReminderByIDUseCase(reminderRepo)
	updateReminderUseCase := application.NewUpdateReminderUseCase(reminderRepo)
	deleteReminderUseCase := application.NewDeleteReminderUseCase(reminderRepo)

	createReminderController := controllers.NewCreateReminderController(createReminderUseCase)
	getAllRemindersController := controllers.NewGetRemindersController(getAllRemindersUseCase) 
	getReminderByIDController := controllers.NewGetReminderByIDController(getReminderByIDUseCase)
	updateReminderController := controllers.NewUpdateReminderController(updateReminderUseCase)
	deleteReminderController := controllers.NewDeleteReminderController(deleteReminderUseCase)

	return routes.NewReminderRoutes(
		createReminderController,
		getAllRemindersController,
		getReminderByIDController,
		updateReminderController,
		deleteReminderController,
	)
}