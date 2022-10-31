package interfaces

import "tasktracker-api/pkg/models"

type ITask interface {
	GetAll(user int) (models.TaskList, error)
	GetTaskById(int, int) (models.Task, error)
	CreateTask(int, models.TaskData) (int, error)
	UpdateTask(int, int, models.TaskData) error
	DeleteTask(int, int) error
}
type IUser interface {
	// GetAll(user int) (models.TaskList, error)
	GetUserById(int) (models.User, error)
	CreateUser(models.UserData) (int, error)
	// UpdateTask(int, int, models.TaskData) error
	// DeleteTask(int, int) error
}
type ITaskService interface {
	GetAll(user int) (models.TaskList, error)
	GetTaskById(int, int) (models.Task, error)
	CreateTask(int, models.TaskData) (int, error)
	UpdateTask(int, int, models.TaskData) error
	DeleteTask(int, int) error
}
type IAuthService interface {
	// GetAll(user int) (models.TaskList, error)
	Login(models.AuthData) (string, error)
	GetUserById(int) (models.User, error)
	CreateUser(models.UserData) (int, error)
	// UpdateTask(int, int, models.TaskData) error
	// DeleteTask(int, int) error
}
