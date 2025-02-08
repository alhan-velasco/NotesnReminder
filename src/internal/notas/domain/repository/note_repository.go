package repository

import "ARQ.HEX/src/internal/notas/domain/entities"

type NoteRepository interface {
	Save(note entities.Note) error
	FindAll() ([]entities.Note, error)
	FindByID(id int) (entities.Note, error)
	Update(note entities.Note) error
	Delete(id int) error
}
