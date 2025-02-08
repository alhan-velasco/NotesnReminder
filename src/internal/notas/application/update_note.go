package application

import (
	"ARQ.HEX/src/internal/notas/domain/entities"
	"ARQ.HEX/src/internal/notas/domain/repository"
)

type UpdateNote struct {
	repo repository.NoteRepository
}

func NewUpdateNoteUseCase(repo repository.NoteRepository) *UpdateNote {
	return &UpdateNote{repo: repo}
}

func (uc *UpdateNote) Execute(note entities.Note) error {

	return uc.repo.Update(note)
}
