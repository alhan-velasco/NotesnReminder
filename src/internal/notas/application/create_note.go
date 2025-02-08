package application

import (
	"ARQ.HEX/src/internal/notas/domain/entities"
	"ARQ.HEX/src/internal/notas/domain/repository"
)

type CreateNoteUseCase struct {
	repo repository.NoteRepository
}

func NewCreateNoteUseCase(repo repository.NoteRepository) *CreateNoteUseCase {
	return &CreateNoteUseCase{repo: repo}
}

func (uc *CreateNoteUseCase) Execute(note entities.Note) error {

	return uc.repo.Save(note)
}
