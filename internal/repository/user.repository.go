package repository

import (
	"database/sql"

	"github.com/Chetan7595/task-manager/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	query := `
	INSERT INTO users (name, email, password, role)
	VALUES (?, ?, ?, ?)
	`
	result, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.Role)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	user.ID = id
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	query := `
	SELECT id, name, email, password, role, created_at
	FROM users WHERE email = ?
	`

	row := r.db.QueryRow(query, email)

	user := &models.User{}
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
