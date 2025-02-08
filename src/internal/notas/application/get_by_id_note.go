package application

import (
	"ARQ.HEX/src/internal/notas/domain/entities"
	"ARQ.HEX/src/internal/notas/domain/repository"
)

type GetNoteByID struct {
	repo repository.NoteRepository
}

func NewGetNoteByIDUseCase(repo repository.NoteRepository) *GetNoteByID {
	return &GetNoteByID{repo: repo}
}

func (uc *GetNoteByID) Execute(id int) (entities.Note, error) {
	return uc.repo.FindByID(id)
}