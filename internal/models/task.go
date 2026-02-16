package models

import "time"

type Task struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	UserID      int64     `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}
