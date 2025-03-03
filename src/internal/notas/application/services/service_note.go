package services

import (
	"ARQ.HEX/src/internal/notas/domain/repository"
	"ARQ.HEX/src/internal/notas/domain/entities"
)

type NoteService struct {
	repo repository.NoteRepository
}

func NewNoteService(repo repository.NoteRepository) *NoteService {
	return &NoteService{repo: repo}
}

func (service *NoteService) GetAllNotes() ([]entities.Note, error) {
	return service.repo.FindAll()
}

func (service *NoteService) GetNoteById(id int) (entities.Note, error) {
	return service.repo.FindByID(id)
}

func (service *NoteService) UpdateNote(note entities.Note) error {
	return service.repo.Update(note)
}

func (service *NoteService) CreateNote(note entities.Note) error {
	return service.repo.Save(note)
}

func (service *NoteService) DeleteNote(id int) error {
	return service.repo.Delete(id)
}
