package service

import (
	"tasktracker-api/pkg/interfaces"
	"tasktracker-api/pkg/repository"
)

type Task interface{ interfaces.ITaskService }
type Auth interface{ interfaces.IAuthService }
type Group interface{ interfaces.IGroupService }

// type Hub interface{ interfaces.IHubService }
type Service struct {
	Task
	Auth
	Group
}

type userCtx string

const ctxKeyUser userCtx = "user"

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Task:  NewTasksService(repo.Tasks, repo.Groups),
		Auth:  NewAuthService(repo.Users, repo.Tasks),
		Group: NewGroupService(repo.Groups, repo.Tasks),
	}
}
