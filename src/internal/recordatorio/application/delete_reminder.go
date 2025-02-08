package application

import "ARQ.HEX/src/internal/recordatorio/domain/repository"

type DeleteReminder struct {
	repo repository.ReminderRepository
}

func NewDeleteReminderUseCase(repo repository.ReminderRepository) *DeleteReminder {
	return &DeleteReminder{repo: repo}
}

func (uc *DeleteReminder) Execute(id int) error {
	return uc.repo.Delete(id)
}
