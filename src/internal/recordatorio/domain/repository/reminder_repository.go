package repository

import "ARQ.HEX/src/internal/recordatorio/domain/entities"

type ReminderRepository interface {
	Create(reminder entities.Reminder) error
	GetAll() ([]entities.Reminder, error)
	GetByID(id int) (entities.Reminder, error)
	Update(reminder entities.Reminder) error
	Delete(id int) error
}