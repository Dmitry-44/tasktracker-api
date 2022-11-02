package router

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"tasktracker-api/pkg/service"

	"github.com/gin-gonic/gin"
)

type Router struct {
	services *service.Service
}

func NewRouter(services *service.Service) *Router {
	return &Router{services: services}
}

func (r *Router) InitRoutes() *gin.Engine {

	router := gin.New()
	router.POST("/login", r.Login)
	router.POST("/logup", r.Logup)
	api := router.Group("/api")
	api.Use(AuthMiddleware(r))
	{
		v1 := api.Group("/v1")
		{
			tasks := v1.Group(("/tasks"))
			{
				tasks.GET("/", r.GetAllTasks)
				tasks.GET("/:id", r.GetTaskById)
				tasks.POST("/", r.CreateTask)
				tasks.PUT("/:id", r.UpdateTask)
				tasks.DELETE("/:id", r.DeleteTask)

			}
			// groups := v1.Group("/groups")
			{
				// groups.GET("/", r.GetAllGroupes)
				// groups.GET("/:id", r.GetGroupById)
				// groups.POST("/", r.CreateGroup)
				// groups.PUT("/:id", r.UpdateGroup)
				// groups.DELETE("/:id", r.DeleteGroup)
			}
		}
	}
	return router
}

type userCtx string

const ctxKeyUser userCtx = "user"

func AuthMiddleware(r *Router) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Print("auth middleware")
		user_id := extractToken(c)
		userId, err := strconv.Atoi(user_id)

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		user, err := r.services.Auth.GetUserById(userId)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		fmt.Printf("user is %v", user)
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), ctxKeyUser, user))
		c.Next()
	}
}

func extractToken(c *gin.Context) string {
	bearToken := c.GetHeader("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
