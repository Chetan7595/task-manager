package repository

import (
	"database/sql"

	"github.com/Chetan7595/task-manager/internal/models"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task *models.Task) error {
	query := `
	INSERT INTO tasks (title, description, status, user_id)
	VALUES (?, ?, ?, ?)
	`

	result, err := r.db.Exec(query,
		task.Title,
		task.Description,
		task.Status,
		task.UserID,
	)

	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	task.ID = id
	return nil
}
