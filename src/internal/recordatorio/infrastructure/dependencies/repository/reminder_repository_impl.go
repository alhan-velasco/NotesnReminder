package repository

import (
	"ARQ.HEX/src/internal/recordatorio/domain/entities"
	"database/sql"
	"fmt"
)

type ReminderRepositoryImpl struct {
	db *sql.DB
}

func NewReminderRepositoryImpl(db *sql.DB) *ReminderRepositoryImpl {
	return &ReminderRepositoryImpl{db: db}
}

func (r *ReminderRepositoryImpl) Create(reminder entities.Reminder) error {
	query := "INSERT INTO reminders (title, description, date) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, reminder.Title, reminder.Description, reminder.DateTime)
	if err != nil {
		return fmt.Errorf("error al guardar recordatorio: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener el ID insertado: %w", err)
	}

	reminder.ID = int32(id)
	return nil
}

func (r *ReminderRepositoryImpl) GetAll() ([]entities.Reminder, error) {
	query := "SELECT id, title, description, date FROM reminders"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener recordatorios: %w", err)
	}
	defer rows.Close()

	var reminders []entities.Reminder
	for rows.Next() {
		var reminder entities.Reminder
		if err := rows.Scan(&reminder.ID, &reminder.Title, &reminder.Description, &reminder.DateTime); err != nil {
			return nil, fmt.Errorf("error al escanear recordatorio: %w", err)
		}
		reminders = append(reminders, reminder)
	}

	return reminders, nil
}

func (r *ReminderRepositoryImpl) GetByID(id int) (entities.Reminder, error) {
	query := "SELECT id, title, description, date FROM reminders WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var reminder entities.Reminder
	if err := row.Scan(&reminder.ID, &reminder.Title, &reminder.Description, &reminder.DateTime); err != nil {
		if err == sql.ErrNoRows {
			return entities.Reminder{}, fmt.Errorf("recordatorio no encontrado")
		}
		return entities.Reminder{}, fmt.Errorf("error al obtener recordatorio: %w", err)
	}

	return reminder, nil
}

func (r *ReminderRepositoryImpl) Update(reminder entities.Reminder) error {
	query := "UPDATE reminders SET title = ?, description = ?, date = ? WHERE id = ?"
	result, err := r.db.Exec(query, reminder.Title, reminder.Description, reminder.DateTime, reminder.ID)
	if err != nil {
		return fmt.Errorf("error al actualizar recordatorio: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("recordatorio no encontrado para actualizar")
	}

	return nil
}

func (r *ReminderRepositoryImpl) Delete(id int) error {
	query := "DELETE FROM reminders WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar recordatorio: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("recordatorio no encontrado para eliminar")
	}

	return nil
}