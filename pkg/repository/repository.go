package repository

import (
	"database/sql"
	"tasktracker-api/pkg/interfaces"
)

type Users interface{ interfaces.IUser }
type Tasks interface{ interfaces.ITask }
type Groups interface{ interfaces.IGroup }

type Repository struct {
	Tasks
	Users
	Groups
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Tasks:  NewTasksRepo(db),
		Users:  NewUsersRepo(db),
		Groups: NewGroupsRepo(db),
	}
}
