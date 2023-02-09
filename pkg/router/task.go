package router

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"tasktracker-api/pkg/hub"
	"tasktracker-api/pkg/models"

	"github.com/gin-gonic/gin"
)

func (r *Router) GetAllTasks(ctx *gin.Context) {
	var params models.TaskGetParams
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		log.Printf("GetAllTask error: %v", err)
	}
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
	data, err := r.services.Task.GetAll(user.Id, params)
	if err != nil {
		ctx.IndentedJSON(
			http.StatusBadGateway,
			gin.H{
				"status":       models.StatusError,
				"data":         "[]",
				"errorMessage": err.Error(),
			})
		return
	}
	ctx.IndentedJSON(
		http.StatusOK,
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
	createdTask, err := r.services.Task.CreateTask(user.Id, task)
	if err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status": models.StatusError,
				"data":   "",
				"error":  err.Error(),
			})
		return
	}
	//send by ws channel
	wsMessage := hub.WSMessage{Entity: "task", Action: "create", Data: createdTask}
	ok = r.hub.SendMessage(user.Id, wsMessage)
	if !ok {
		fmt.Print("WS Send Message error")
	}
	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{
			"status":       models.StatusSuccess,
			"data":         createdTask,
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
	updatedTask, err := r.services.Task.UpdateTask(user.Id, id, task)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
	//send by ws channel
	wsMessage := hub.WSMessage{Entity: "task", Action: "update", Data: updatedTask}
	ok = r.hub.SendMessage(user.Id, wsMessage)
	if !ok {
		fmt.Print("WS Send Message error")
	}
	ctx.IndentedJSON(
		http.StatusOK,
		models.ServerResponse{
			Status: "ok",
			Data:   fmt.Sprintf("%v", updatedTask),
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
	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{
			"status":       models.StatusSuccess,
			"data":         taskId,
			"errorMessage": "",
		},
	)
}
