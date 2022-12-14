package router

import (
	"fmt"
	"net/http"
	"tasktracker-api/pkg/models"

	"github.com/gin-gonic/gin"
)

func (r *Router) GetAllGroupes(ctx *gin.Context) {
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
	data, err := r.services.Group.GetAll(user.Id)
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

func (r *Router) GetGroupById(ctx *gin.Context) {
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
	group, err := r.services.Group.GetGroupById(ctx, user.Id)
	if err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status": models.StatusError,
				"data":   group,
				"error":  err.Error()})
		return
	}
	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{
			"status":       models.StatusSuccess,
			"data":         group,
			"errorMessage": "",
		},
	)
}

func (r *Router) CreateGroup(ctx *gin.Context) {
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
	group := models.GroupData{}
	if err := ctx.BindJSON(&group); err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status": models.StatusError,
				"data":   "",
				"error":  fmt.Sprintf("Server error: %v", err.Error()),
			})
		return
	}
	if group.Name == nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status": models.StatusError,
				"data":   "",
				"error":  "Name is required",
			})
		return
	}
	id, err := r.services.Group.CreateGroup(user.Id, group)
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

func (r *Router) DeleteGroup(ctx *gin.Context) {
	user, ok := ctx.Request.Context().Value(ctxKeyUser).(models.User)
	if !ok {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status": models.StatusError,
				"data":   "",
				"error":  "server error",
			})
		return
	}
	err := r.services.Group.DeleteGroup(ctx, user.Id)
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
	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{
			"status":       models.StatusSuccess,
			"data":         "",
			"errorMessage": "",
		},
	)
}
func (r *Router) GetTasksByGroupId(ctx *gin.Context) {
	user, ok := ctx.Request.Context().Value(ctxKeyUser).(models.User)
	if !ok {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status": models.StatusError,
				"data":   "",
				"error":  "server error",
			})
		return
	}
	tasks, err := r.services.Group.GetTasksByGroupId(ctx, user.Id)
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
	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{
			"status":       models.StatusSuccess,
			"data":         tasks,
			"errorMessage": "",
		},
	)
}
