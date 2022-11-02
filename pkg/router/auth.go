package router

import (
	"fmt"
	"net/http"
	"tasktracker-api/pkg/models"

	"github.com/gin-gonic/gin"
)

func (r *Router) Login(ctx *gin.Context) {
	user := models.AuthData{}
	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status": models.StatusError,
				"data":   "",
				"error":  fmt.Sprintf("Server error: %v", err.Error()),
			})
		return
	}
	jwtToken, err := r.services.Auth.Login(user)
	if err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status":       models.StatusError,
				"token":        "",
				"errorMessage": err.Error(),
			})
		return
	}
	// ctx.SetCookie("tasktrackerToken", token, 3600, "/", "/", true, false)
	ctx.IndentedJSON(
		http.StatusCreated,
		gin.H{
			"status":       models.StatusSuccess,
			"token":        jwtToken,
			"errorMessage": "",
		},
	)
}

func (r *Router) Logup(ctx *gin.Context) {
	user := models.UserData{}
	fmt.Printf("user in router %v", user)
	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status": models.StatusError,
				"data":   "",
				"error":  fmt.Sprintf("Server error: %v", err.Error()),
			})
		return
	}
	jwtToken, err := r.services.Auth.Logup(user)
	if err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status":       models.StatusError,
				"token":        "",
				"errorMessage": err.Error(),
			})
		return
	}
	// ctx.SetCookie("tasktrackerToken", token, 3600, "/", "/", true, false)
	ctx.IndentedJSON(
		http.StatusCreated,
		gin.H{
			"status":       models.StatusSuccess,
			"token":        jwtToken,
			"errorMessage": "",
		},
	)
}
