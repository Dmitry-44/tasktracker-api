package router

import (
	"fmt"
	"net/http"
	"strconv"
	"tasktracker-api/pkg/models"

	"github.com/gin-gonic/gin"
)

func (r *Router) GetAllTasks(ctx *gin.Context) {
	user, userExist := ctx.Get("user")
	if userExist != true {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"data": "Unauthorized user"})
		return
	}
	fmt.Printf("context is : %v", user)
	data, err := r.services.Tasks.GetAll(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Server Error"})
		return
	}
	ctx.IndentedJSON(
		http.StatusCreated,
		gin.H{"data": data},
	)
}

func (r *Router) CreateTask(ctx *gin.Context) {
	// user:=models.User{Id:1}
	fmt.Printf("context is : %v", ctx)
	task := models.TaskData{}
	if err := ctx.BindJSON(&task); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Server error"})
		return
	}
	if task.Title == nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Empty title field"})
		return
	}
	id, err := r.services.Tasks.CreateTask(task)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
	ctx.IndentedJSON(
		http.StatusCreated,
		gin.H{"data": id},
	)
}

func (r *Router) GetTaskById(ctx *gin.Context) {
	taskIdStr := ctx.Param("id")
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Server error"})
		return
	}
	task, err := r.services.Tasks.GetTaskById(taskId)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"data": err.Error()})
		return
	}
	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{"data": task},
	)
}

func (r *Router) UpdateTask(ctx *gin.Context) {
	task := models.TaskData{}
	idString := ctx.Param("id")
	if err := ctx.BindJSON(&task); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Server error"})
		return
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.IndentedJSON(http.StatusOK, gin.H{"data": "Server error"})
		return
	}
	err = r.services.Tasks.UpdateTask(id, task)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
	ctx.IndentedJSON(
		http.StatusOK,
		models.ServerResponse{
			Status: "ok",
			Data:   idString,
		},
	)
}

func (r *Router) DeleteTask(ctx *gin.Context) {
	taskIdStr := ctx.Param("id")
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Server error"})
		return
	}
	task, err := r.services.Tasks.GetTaskById(taskId)
	if err != nil {
		errorMessage := fmt.Sprintf("Data with id = %v not found", taskId)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"dataa": errorMessage})
		return
	}
	if (models.Task{} == task) {
		errorMessage := fmt.Sprintf("Data with id = %v not found!", taskId)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"data": errorMessage})
		return
	}
	err = r.services.Tasks.DeleteTask(taskId)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
	respString := fmt.Sprintf("data with id=%v was deleted", taskId)
	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{"data": respString},
	)
}
