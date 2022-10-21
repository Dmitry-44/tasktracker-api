package service

import (
	"tasktracker-api/pkg/models"
	"tasktracker-api/pkg/repository"
)

type Tasks interface {
	GetAll() (models.TaskList, error)
	GetTaskById(int) (models.Task, error)
	CreateTask(models.TaskData) (int, error)
	UpdateTask(int, models.TaskData) (int, error)
	DeleteTask(int) error
}

type Service struct {
	Tasks
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Tasks: NewTaskService(repo.Tasks),
	}
}
