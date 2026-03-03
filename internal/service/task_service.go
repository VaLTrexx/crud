package service

import (
	"context"
	"errors"

	"github.com/VaLTrexx/crud/internal/models"
	"github.com/VaLTrexx/crud/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(ctx context.Context, task *models.Task) error {

	if task.Title == "" {
		return errors.New("title cannot be empty")
	}

	if task.Priority < 1 || task.Priority > 3 {
		return errors.New("priority must be between 1 and 3")
	}

	return s.repo.CreateTask(ctx, task)
}

func (s *TaskService) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	return s.repo.GetAllTasks(ctx)
}
