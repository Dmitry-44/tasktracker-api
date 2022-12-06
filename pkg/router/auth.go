package router

import (
	"fmt"
	"net/http"
	"strconv"
	"tasktracker-api/pkg/models"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	TokenLifeTime = int(24 * time.Hour)
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
	ctx.SetCookie("Bearer", jwtToken, TokenLifeTime, "/", ctx.Request.Header.Get("Origin"), false, true)
	ctx.IndentedJSON(
		http.StatusOK,
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

func (r *Router) Auth(ctx *gin.Context) {
	token := extractTokenFromHeader(ctx)
	if len(token) == 0 {
		ctx.IndentedJSON(
			http.StatusUnauthorized,
			gin.H{
				"status":       models.StatusError,
				"data":         "",
				"errorMessage": "unauthorized",
			},
		)
		return
	}
	claims, ok := GetClaimsFromToken(token)
	if ok != nil {
		ctx.IndentedJSON(
			http.StatusUnauthorized,
			gin.H{
				"status":       models.StatusError,
				"data":         "",
				"errorMessage": "token error",
			},
		)
		return
	}
	userIDString := claims["sub"].(string)
	userId, err := strconv.Atoi(userIDString)
	if err != nil {
		ctx.IndentedJSON(
			http.StatusUnauthorized,
			gin.H{
				"status":       models.StatusError,
				"data":         "",
				"errorMessage": err.Error(),
			},
		)
		return
	}
	user, err := r.services.Auth.GetUserById(userId)
	if err != nil {
		ctx.IndentedJSON(
			http.StatusUnauthorized,
			gin.H{
				"status":       models.StatusError,
				"data":         "",
				"errorMessage": err.Error(),
			},
		)
		return
	}
	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{
			"status":       models.StatusSuccess,
			"data":         user,
			"errorMessage": "",
		},
	)
}
