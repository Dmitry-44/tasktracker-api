package service

import (
	"fmt"
	"tasktracker-api/pkg/models"
	"tasktracker-api/pkg/repository"
)

type TaskService struct {
	repo repository.Tasks
}

func NewTaskService(repo repository.Tasks) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAll(userId int) ([]models.Task, error) {
	fmt.Println("GetAll")
	return s.repo.GetAll(userId)
}
