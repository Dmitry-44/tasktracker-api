package interfaces

import (
	"tasktracker-api/pkg/models"

	"github.com/gin-gonic/gin"
)

type ITask interface {
	GetAll(user int) (models.TaskList, error)
	GetTaskById(int, int) (models.Task, error)
	CreateTask(int, models.TaskData) (models.Task, error)
	UpdateTask(int, int, models.TaskData) error
	DeleteTask(int, int) error
	GetTasksByGroupId(id int) (models.TaskList, error)
}
type IUser interface {
	// GetAll(user int) (models.TaskList, error)
	GetUserById(int) (models.User, error)
	GetUserByLogin(string) (models.User, error)
	CreateUser(models.UserData) (int, error)
	// UpdateTask(int, int, models.TaskData) error
	// DeleteTask(int, int) error
}
type ITaskService interface {
	GetAll(user int) (models.TaskList, error)
	GetTaskById(int, int) (models.Task, error)
	CreateTask(int, models.TaskData) (models.Task, error)
	UpdateTask(int, int, models.TaskData) error
	DeleteTask(int, int) error
}
type IAuthService interface {
	// GetAll(user int) (models.TaskList, error)
	Login(models.AuthData) (string, models.User, error)
	Logup(models.UserData) (string, error)
	GetUserById(int) (models.User, error)
	CreateUser(models.UserData) (int, error)
	// UpdateTask(int, int, models.TaskData) error
	// DeleteTask(int, int) error
}

type IGroupService interface {
	GetAll(int) (models.GroupList, error)
	GetGroupById(ctx *gin.Context, userId int) (models.Group, error)
	CreateGroup(int, models.GroupData) (int, error)
	DeleteGroup(ctx *gin.Context, userId int) error
	GetTasksByGroupId(ctx *gin.Context, userId int) (models.TaskList, error)
}
type IGroup interface {
	GetAll(int) (models.GroupList, error)
	GetGroupById(int) (models.Group, error)
	CreateGroup(int, models.GroupData) (int, error)
	// SetUserGroup(int, int) error
	DeleteGroupById(int) error
}
