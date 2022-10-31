package router

import (
	"fmt"
	"net/http"
	"strconv"
	"tasktracker-api/pkg/models"

	"github.com/gin-gonic/gin"
)

func (r *Router) GetAllTasks(ctx *gin.Context) {
	user, ok := ctx.Request.Context().Value(ctxKeyUser).(models.User)
	if !ok {
		ctx.IndentedJSON(
			http.StatusUnauthorized,
			gin.H{
				"status":       models.StatusSuccess,
				"data":         "[]",
				"errorMessage": "Unauthorized user",
			})
		return
	}
	data, err := r.services.Task.GetAll(user.Id)
	if err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status":       models.StatusError,
				"data":         "[]",
				"errorMessage": err.Error(),
			})
		return
	}
	ctx.IndentedJSON(
		http.StatusCreated,
		gin.H{
			"status":       models.StatusSuccess,
			"data":         data,
			"errorMessage": "",
		},
	)
}

func (r *Router) GetTaskById(ctx *gin.Context) {
	user, ok := ctx.Request.Context().Value(ctxKeyUser).(models.User)
	if !ok {
		ctx.IndentedJSON(
			http.StatusUnauthorized,
			gin.H{
				"status":       models.StatusSuccess,
				"data":         "[]",
				"errorMessage": "Unauthorized user",
			})
		return
	}
	taskIdStr := ctx.Param("id")
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status": models.StatusError,
				"data":   "",
				"error":  fmt.Sprintf("Server error: %v", err.Error()),
			})
		return
	}
	task, err := r.services.Task.GetTaskById(user.Id, taskId)
	if err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status": models.StatusError,
				"data":   task,
				"error":  err.Error()})
		return
	}
	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{
			"status":       models.StatusSuccess,
			"data":         task,
			"errorMessage": "",
		},
	)
}

func (r *Router) CreateTask(ctx *gin.Context) {
	user, ok := ctx.Request.Context().Value(ctxKeyUser).(models.User)
	if !ok {
		ctx.IndentedJSON(
			http.StatusUnauthorized,
			gin.H{
				"status":       models.StatusSuccess,
				"data":         "",
				"errorMessage": "Unauthorized user",
			})
		return
	}
	task := models.TaskData{}
	if err := ctx.BindJSON(&task); err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status": models.StatusError,
				"data":   "",
				"error":  fmt.Sprintf("Server error: %v", err.Error()),
			})
		return
	}
	if task.Title == nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status": models.StatusError,
				"data":   "",
				"error":  "Title is required",
			})
		return
	}
	id, err := r.services.Task.CreateTask(user.Id, task)
	if err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status": models.StatusError,
				"data":   id,
				"error":  err.Error(),
			})
		return
	}
	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{
			"status":       models.StatusSuccess,
			"data":         id,
			"errorMessage": "",
		},
	)
}

func (r *Router) UpdateTask(ctx *gin.Context) {
	user, ok := ctx.Request.Context().Value(ctxKeyUser).(models.User)
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
	err = r.services.Task.UpdateTask(user.Id, id, task)
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
	user, ok := ctx.Request.Context().Value(ctxKeyUser).(models.User)
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
	err = r.services.Task.DeleteTask(user.Id, taskId)
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
