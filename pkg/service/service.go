package service

import (
	"tasktracker-api/pkg/interfaces"
	"tasktracker-api/pkg/repository"
)

type Task interface{ interfaces.ITaskService }
type Auth interface{ interfaces.IAuthService }
type Group interface{ interfaces.IGroupService }
type Service struct {
	Task
	Auth
	Group
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Task:  NewTasksService(repo.Tasks),
		Auth:  NewAuthService(repo.Users, repo.Tasks),
		Group: NewGroupService(repo.Groups),
	}
}
