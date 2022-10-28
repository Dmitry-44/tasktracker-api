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
	router.Use(AuthMiddleware())
	api := router.Group("/api")
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

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := extractToken(c)
		user, err := strconv.Atoi(user_id)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		fmt.Print("here")
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "user_id", user))
		c.Next()
	}
}

func extractToken(c *gin.Context) string {
	bearToken := c.GetHeader("Authorization")
	strArr := strings.Split(bearToken, " ")
	fmt.Printf("auth string is : %s", bearToken)
	fmt.Printf("auth arr is : %v", strArr)
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
