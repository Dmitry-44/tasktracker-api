package repository

import (
	"database/sql"
	"tasktracker-api/pkg/models"
)

type Tasks interface {
	GetAll(userId int) ([]models.Task, error)
}
type Repository struct {
	Tasks
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Tasks: NewTaskRepo(db),
	}
}
