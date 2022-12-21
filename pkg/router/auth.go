package router

import (
	"fmt"
	"net/http"
	"strconv"
	"tasktracker-api/pkg/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	TokenLifeTime = int((time.Duration(24) * time.Hour) / 1000000000) //sec
)

func (r *Router) Login(ctx *gin.Context) {
	user := models.AuthData{}
	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status":       models.StatusError,
				"token":        "",
				"user":         "",
				"errorMessage": fmt.Sprintf("Server error: %v", err.Error()),
			})
		return
	}
	jwtToken, userFromDb, err := r.services.Auth.Login(user)
	if err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status":       models.StatusError,
				"token":        "",
				"user":         userFromDb,
				"errorMessage": err.Error(),
			})
		return
	}
	ctx.SetCookie(viper.GetString("BearerCookieName"), jwtToken, TokenLifeTime, "/", ctx.Request.Header.Get("Origin"), true, false)
	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{
			"status":       models.StatusSuccess,
			"token":        jwtToken,
			"user":         userFromDb,
			"errorMessage": "",
		},
	)
}

func (r *Router) Logup(ctx *gin.Context) {
	user := models.UserData{}
	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status":       models.StatusError,
				"token":        "",
				"user":         "",
				"errorMessage": fmt.Sprintf("Server error: %v", err.Error()),
			})
		return
	}
	jwtToken, userFromDB, err := r.services.Auth.Logup(user)
	if err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status":       models.StatusError,
				"token":        "",
				"user":         "",
				"errorMessage": err.Error(),
			})
		return
	}
	ctx.SetCookie("Bearer", jwtToken, TokenLifeTime, "/", ctx.Request.Header.Get("Origin"), false, false)
	ctx.IndentedJSON(
		http.StatusCreated,
		gin.H{
			"status":       models.StatusSuccess,
			"token":        jwtToken,
			"user":         userFromDB,
			"errorMessage": "",
		},
	)
}

func (r *Router) Auth(ctx *gin.Context) {
	token := r.extractTokenFromHeader(ctx)
	if len(token) == 0 {
		ctx.IndentedJSON(
			http.StatusUnauthorized,
			gin.H{
				"status":       models.StatusError,
				"user":         "",
				"errorMessage": "unauthorized",
			},
		)
		return
	}
	claims, ok := r.GetClaimsFromToken(token)
	if ok != nil {
		ctx.IndentedJSON(
			http.StatusUnauthorized,
			gin.H{
				"status":       models.StatusError,
				"user":         "",
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
				"user":         "",
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
				"user":         "",
				"errorMessage": err.Error(),
			},
		)
		return
	}
	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{
			"status":       models.StatusSuccess,
			"user":         user,
			"errorMessage": "",
		},
	)
}
