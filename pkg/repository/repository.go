package repository

import "tasktracker-api/pkg/models"

type Tasks interface {
	GetAll(userId int) ([]models.Task, error)
}
type Repository struct {
	Tasks
}

func NewRepository() *Repository {
	return &Repository{
		Tasks: NewTaskStore(),
	}
}
