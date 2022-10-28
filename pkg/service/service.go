package service

import (
	"tasktracker-api/pkg/models"
	"tasktracker-api/pkg/repository"
)

type Tasks interface {
	GetAll(user int) (models.TaskList, error)
	GetTaskById(int, int) (models.Task, error)
	CreateTask(int, models.TaskData) (int, error)
	UpdateTask(int, models.TaskData) error
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
