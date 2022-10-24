package repository

import (
	"database/sql"
	"tasktracker-api/pkg/models"
)

type Tasks interface {
	GetAll() (models.TaskList, error)
	GetTaskById(int) (models.Task, error)
	CreateTask(models.TaskData) (int, error)
	UpdateTask(int, models.TaskData) error
	DeleteTask(int) error
}
type Repository struct {
	Tasks
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Tasks: NewTaskRepo(db),
	}
}
