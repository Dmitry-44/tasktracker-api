package service

import (
	"tasktracker-api/pkg/models"
	"tasktracker-api/pkg/repository"
)

type AuthService struct {
	usersRepo repository.Users
	tasksRepo repository.Tasks
}

func NewAuthService(usersRepo repository.Users, tasksRepo repository.Tasks) *AuthService {
	return &AuthService{
		usersRepo: usersRepo,
		tasksRepo: tasksRepo,
	}
}

// func (s *AuthService) GetAll(user int) (models.TaskList, error) {
// 	return s.repo.GetAll(user)
// }
func (s *AuthService) GetUserById(userId int) (models.User, error) {
	return s.usersRepo.GetUserById(userId)
}

// func (s *AuthService) CreateUser(task models.UserData) (int, error) {
// 	return s.usersRepo.CreateUser(task)
// }

// func (s *AuthService) UpdateTask(user int, taskId int, task models.TaskData) error {
// 	return s.repo.UpdateTask(user, taskId, task)
// }
// func (s *AuthService) DeleteTask(user int, taskId int) error {
// 	return s.repo.DeleteTask(user, taskId)
// }
