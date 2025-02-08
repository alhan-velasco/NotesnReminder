package application

import (
	"ARQ.HEX/src/internal/notas/domain/entities"
	"ARQ.HEX/src/internal/notas/domain/repository"
)

type GetNotes struct {
	repo repository.NoteRepository
}

func NewGetNotesUseCase(repo repository.NoteRepository) *GetNotes {
	return &GetNotes{repo: repo}
}

func (uc *GetNotes) Execute() ([]entities.Note, error) {
	return uc.repo.FindAll()
}
