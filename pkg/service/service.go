package service

import (
	"tasktracker-api/pkg/interfaces"
	"tasktracker-api/pkg/repository"
)

type IUser interface{ interfaces.IUser }
type ITask interface{ interfaces.ITask }
type Service struct {
	ITask
	IUser
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		ITask: NewTasksService(repo.ITask),
		IUser: NewAuthService(repo.IUser, repo.ITask),
	}
}
