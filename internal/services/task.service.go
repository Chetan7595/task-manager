package services

import (
	"github.com/Chetan7595/task-manager/internal/models"
	"github.com/Chetan7595/task-manager/internal/repository"
)

type TaskService struct {
	taskRepo *repository.TaskRepository
}

func NewTaskService(taskRepo *repository.TaskRepository) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) Create(title, desc string, userID int64) error {

	task := &models.Task{
		Title:       title,
		Description: desc,
		Status:      "PENDING",
		UserID:      userID,
	}

	return s.taskRepo.Create(task)
}
