package repository

import (
	"database/sql"
	"tasktracker-api/pkg/models"
)

type Tasks interface {
	GetAll(user int) (models.TaskList, error)
	GetTaskById(int, int) (models.Task, error)
	CreateTask(int, models.TaskData) (int, error)
	UpdateTask(int, int, models.TaskData) error
	DeleteTask(int, int) error
}
type Repository struct {
	Tasks
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Tasks: NewTaskRepo(db),
	}
}
