package application

import "ARQ.HEX/src/internal/notas/domain/repository"

type DeleteNote struct {
	repo repository.NoteRepository
}

func NewDeleteNoteUseCase(repo repository.NoteRepository) *DeleteNote {
	return &DeleteNote{repo: repo}
}

func (uc *DeleteNote) Execute(id int) error {
	return uc.repo.Delete(id)
}
