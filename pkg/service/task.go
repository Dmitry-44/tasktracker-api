package service

import (
	"errors"
	"tasktracker-api/pkg/models"
	"tasktracker-api/pkg/repository"
)

type TasksService struct {
	tasksRepo repository.Tasks
	groupRepo repository.Groups
}

func NewTasksService(tasksRepo repository.Tasks, groupRepo repository.Groups) *TasksService {
	return &TasksService{
		tasksRepo: tasksRepo,
		groupRepo: groupRepo,
	}
}

func (s *TasksService) GetAll(user int, params models.TaskGetParams) (models.TaskList, error) {
	if params.GroupId != nil {
		taskList := models.TaskList{}
		taskList.Tasks = make([]models.Task, 0)
		userIsMember, err := s.groupRepo.IsUserInGroup(user, *params.GroupId)
		if err != nil {
			return taskList, errors.New("permission denied")
		}
		if userIsMember {
			return s.tasksRepo.GetTasksByGroupId(*params.GroupId)
		} else {
			return taskList, errors.New("permission denied")
		}
	}
	return s.tasksRepo.GetAll(user)
}
func (s *TasksService) GetTaskById(user int, taskId int) (models.Task, error) {
	return s.tasksRepo.GetTaskById(user, taskId)
}
func (s *TasksService) CreateTask(user int, task models.TaskData) (models.Task, error) {
	return s.tasksRepo.CreateTask(user, task)
}
func (s *TasksService) UpdateTask(user int, taskId int, task models.TaskData) (models.Task, error) {
	return s.tasksRepo.UpdateTask(user, taskId, task)
}
func (s *TasksService) DeleteTask(user int, taskId int) error {
	return s.tasksRepo.DeleteTask(user, taskId)
}
