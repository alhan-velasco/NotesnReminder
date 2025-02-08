package application

import (
    "ARQ.HEX/src/internal/recordatorio/domain/entities"
    "ARQ.HEX/src/internal/recordatorio/domain/repository"
)

type GetReminder struct {
	repo repository.ReminderRepository
}

func NewGetReminderUseCase(repo repository.ReminderRepository) *GetReminder {
	return &GetReminder{repo: repo}
}

func (uc *GetReminder) Execute() ([]entities.Reminder, error) {
	return uc.repo.GetAll()
}
