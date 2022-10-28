package service

import (
	"tasktracker-api/pkg/models"
	"tasktracker-api/pkg/repository"
)

type TaskService struct {
	repo repository.Tasks
}

func NewTaskService(repo repository.Tasks) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAll(user int) (models.TaskList, error) {
	return s.repo.GetAll(user)
}
func (s *TaskService) GetTaskById(user int, taskId int) (models.Task, error) {
	return s.repo.GetTaskById(user, taskId)
}
func (s *TaskService) CreateTask(user int, task models.TaskData) (int, error) {
	return s.repo.CreateTask(user, task)
}
func (s *TaskService) UpdateTask(user int, taskId int, task models.TaskData) error {
	return s.repo.UpdateTask(user, taskId, task)
}
func (s *TaskService) DeleteTask(user int, taskId int) error {
	return s.repo.DeleteTask(user, taskId)
}
