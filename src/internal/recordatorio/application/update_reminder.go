package application

import (
    "ARQ.HEX/src/internal/recordatorio/domain/entities"
    "ARQ.HEX/src/internal/recordatorio/domain/repository"
)

type UpdateReminder struct {
	repo repository.ReminderRepository
}

func NewUpdateReminderUseCase(repo repository.ReminderRepository) *UpdateReminder {
	return &UpdateReminder{repo: repo}
}

func (uc *UpdateReminder) Execute(reminder entities.Reminder) error {

	return uc.repo.Update(reminder)
}
