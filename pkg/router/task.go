package router

import (
	"fmt"
	"net/http"
	"strconv"
	"tasktracker-api/pkg/models"

	"github.com/gin-gonic/gin"
)

func (r *Router) GetAllTasks(ctx *gin.Context) {
	user, ok := ctx.Request.Context().Value("user_id").(int)
	if !ok {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"data": "Unauthorized user"})
		return
	}
	data, err := r.services.Tasks.GetAll(user)
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
	user, ok := ctx.Request.Context().Value("user_id").(int)
	if !ok {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"data": "Unauthorized user"})
		return
	}
	task := models.TaskData{}
	if err := ctx.BindJSON(&task); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Server error"})
		return
	}
	if task.Title == nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Empty title field"})
		return
	}
	id, err := r.services.Tasks.CreateTask(user, task)
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
	user, ok := ctx.Request.Context().Value("user_id").(int)
	if !ok {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"data": "Unauthorized user"})
		return
	}
	taskIdStr := ctx.Param("id")
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Server error"})
		return
	}
	task, err := r.services.Tasks.GetTaskById(user, taskId)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"data": task, "error": err.Error()})
		return
	}
	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{"data": task},
	)
}

func (r *Router) UpdateTask(ctx *gin.Context) {
	user, ok := ctx.Request.Context().Value("user_id").(int)
	if !ok {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"data": "Unauthorized user"})
		return
	}
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
	err = r.services.Tasks.UpdateTask(user, id, task)
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
	user, ok := ctx.Request.Context().Value("user_id").(int)
	if !ok {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"data": "Unauthorized user"})
		return
	}
	taskIdStr := ctx.Param("id")
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Server error"})
		return
	}
	task, err := r.services.Tasks.GetTaskById(user, taskId)
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
	err = r.services.Tasks.DeleteTask(user, taskId)
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
