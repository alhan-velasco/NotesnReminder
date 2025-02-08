package application

import (
    "ARQ.HEX/src/internal/recordatorio/domain/entities"
    "ARQ.HEX/src/internal/recordatorio/domain/repository"
)

type GetReminderByID struct {
	repo repository.ReminderRepository
}

func NewGetReminderByIDUseCase(repo repository.ReminderRepository) *GetReminderByID {
	return &GetReminderByID{repo: repo}
}

func (uc *GetReminderByID) Execute(id int) (entities.Reminder, error) {
	return uc.repo.GetByID(id)
}