package repository

import (
	"database/sql"
	"tasktracker-api/pkg/interfaces"
)

type IUser interface{ interfaces.IUser }
type ITask interface{ interfaces.ITask }

type Repository struct {
	ITask
	IUser
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		ITask: NewTasksRepo(db),
		IUser: NewUsersRepo(db),
	}
}
