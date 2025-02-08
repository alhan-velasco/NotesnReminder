package application

import (
	"ARQ.HEX/src/internal/recordatorio/domain/entities"
	"ARQ.HEX/src/internal/recordatorio/domain/repository"
)

type CreateReminder struct {
	repo repository.ReminderRepository
}

func NewCreateReminder(repo repository.ReminderRepository) *CreateReminder {
	return &CreateReminder{repo: repo}
}

func (uc *CreateReminder) Execute(reminder entities.Reminder) error {

	return uc.repo.Create(reminder)
}
