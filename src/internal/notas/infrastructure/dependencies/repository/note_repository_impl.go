package repository

import (
	"database/sql"
	"fmt"

	"ARQ.HEX/src/internal/notas/domain/entities"
)

type NoteRepositoryImpl struct {
	db *sql.DB
}

func NewNoteRepositoryImpl(db *sql.DB) *NoteRepositoryImpl {
	return &NoteRepositoryImpl{db: db}
}

func (r *NoteRepositoryImpl) Save(note entities.Note) error {
	query := "INSERT INTO notas (title, content) VALUES (?, ?)"
	result, err := r.db.Exec(query, note.Title, note.Content)
	if err != nil {
		return fmt.Errorf("error al guardar la nota: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener el ID insertado: %w", err)
	}

	note.ID = int(id)
	return nil
}

func (r *NoteRepositoryImpl) FindAll() ([]entities.Note, error) {
	query := "SELECT id, title, content FROM notas"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener notas: %w", err)
	}
	defer rows.Close()

	var notes []entities.Note
	for rows.Next() {
		var note entities.Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Content); err != nil {
			return nil, fmt.Errorf("error al escanear nota: %w", err)
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (r *NoteRepositoryImpl) FindByID(id int) (entities.Note, error) {
	query := "SELECT id, title, content FROM notas WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var note entities.Note
	if err := row.Scan(&note.ID, &note.Title, &note.Content); err != nil {
		if err == sql.ErrNoRows {
			return entities.Note{}, fmt.Errorf("nota no encontrada")
		}
		return entities.Note{}, fmt.Errorf("error al obtener nota: %w", err)
	}
	return note, nil
}

func (r *NoteRepositoryImpl) Update(note entities.Note) error {
	query := "UPDATE notas SET title = ?, content = ? WHERE id = ?"
	result, err := r.db.Exec(query, note.Title, note.Content, note.ID)
	if err != nil {
		return fmt.Errorf("error al actualizar nota: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("nota no encontrada para actualizar")
	}

	return nil
}

func (r *NoteRepositoryImpl) Delete(id int) error {
	query := "DELETE FROM notas WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar nota: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("nota no encontrada para eliminar")
	}

	return nil
}
