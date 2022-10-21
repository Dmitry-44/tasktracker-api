package router

import (
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
	api := router.Group("/api")
	{
		tasks := api.Group(("/tasks"))
		{
			tasks.GET("/", r.GetAll)
			tasks.GET("/:id", r.GetTaskById)
			tasks.POST("/", r.CreateTask)
			tasks.PUT("/:id", r.UpdateTask)
			tasks.DELETE("/:id", r.DeleteTask)

		}
	}
	return router
}
