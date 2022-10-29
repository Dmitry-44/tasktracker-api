package service

import (
	"tasktracker-api/pkg/models"
	"tasktracker-api/pkg/repository"
)

type TasksService struct {
	tasksRepo repository.ITask
}

func NewTasksService(tasksRepo repository.ITask) *TasksService {
	return &TasksService{tasksRepo: tasksRepo}
}

func (s *TasksService) GetAll(user int) (models.TaskList, error) {
	return s.tasksRepo.GetAll(user)
}
func (s *TasksService) GetTaskById(user int, taskId int) (models.Task, error) {
	return s.tasksRepo.GetTaskById(user, taskId)
}
func (s *TasksService) CreateTask(user int, task models.TaskData) (int, error) {
	return s.tasksRepo.CreateTask(user, task)
}
func (s *TasksService) UpdateTask(user int, taskId int, task models.TaskData) error {
	return s.tasksRepo.UpdateTask(user, taskId, task)
}
func (s *TasksService) DeleteTask(user int, taskId int) error {
	return s.tasksRepo.DeleteTask(user, taskId)
}
