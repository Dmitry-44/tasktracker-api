package repository

import "tasktracker-api/pkg/models"

type TaskStore struct {
	Tasks []models.Task
}

func NewTaskStore() *TaskStore {
	var tasks []models.Task
	task1 := models.Task{Id: 1, Title: "Задача 1"}
	task2 := models.Task{Id: 2, Title: "Задача 2"}
	tasks = append(tasks, task1, task2)
	return &TaskStore{
		Tasks: tasks,
	}
}

func (s *TaskStore) GetAll(userId int) ([]models.Task, error) {
	return s.Tasks, nil
}
