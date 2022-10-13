package service

import (
	"tasktracker-api/pkg/models"
	"tasktracker-api/pkg/repository"
)

type Tasks interface {
	GetAll(userId int) ([]models.Task, error)
}

type Service struct {
	Tasks
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Tasks: NewTaskService(repo.Tasks),
	}
}
