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

func (s *TaskService) GetAll() (models.TaskList, error) {
	return s.repo.GetAll()
}
func (s *TaskService) GetTaskById(taskId int) (models.Task, error) {
	return s.repo.GetTaskById(taskId)
}
func (s *TaskService) CreateTask(task models.TaskData) (int, error) {
	return s.repo.CreateTask(task)
}
func (s *TaskService) UpdateTask(taskId int, task models.TaskData) (int, error) {
	return s.repo.UpdateTask(taskId, task)
}
func (s *TaskService) DeleteTask(taskId int) error {
	return s.repo.DeleteTask(taskId)
}
