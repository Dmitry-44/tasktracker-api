package service

import (
	"tasktracker-api/pkg/models"
	"tasktracker-api/pkg/repository"
)

type GroupsService struct {
	groupsRepo repository.Groups
}

func NewGroupService(groupsRepo repository.Groups) *GroupsService {
	return &GroupsService{groupsRepo: groupsRepo}
}

//	func (s *TasksService) GetAll(user int) (models.TaskList, error) {
//		return s.tasksRepo.GetAll(user)
//	}
//
//	func (s *TasksService) GetTaskById(user int, taskId int) (models.Task, error) {
//		return s.tasksRepo.GetTaskById(user, taskId)
//	}
func (s *GroupsService) CreateGroup(user int, group models.GroupData) (int, error) {
	createdGroupId, err := s.groupsRepo.CreateGroup(user, group)
	if err != nil {
		return 0, err
	}
	err = s.groupsRepo.SetUserGroup(user, createdGroupId)
	if err != nil {
		return 0, err
	}
	return createdGroupId, nil
}

// func (s *TasksService) UpdateTask(user int, taskId int, task models.TaskData) error {
// 	return s.tasksRepo.UpdateTask(user, taskId, task)
// }
// func (s *TasksService) DeleteTask(user int, taskId int) error {
// 	return s.tasksRepo.DeleteTask(user, taskId)
// }
