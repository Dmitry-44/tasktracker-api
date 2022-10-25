package service

import (
	"tasktracker-api/pkg/models"
	"tasktracker-api/pkg/repository"

	"github.com/gin-gonic/gin"
)

type Tasks interface {
	GetAll(*gin.Context) (models.TaskList, error)
	GetTaskById(int) (models.Task, error)
	CreateTask(models.TaskData) (int, error)
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
