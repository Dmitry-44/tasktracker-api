package main

import (
	"tasktracker-api/pkg/repository"
	router "tasktracker-api/pkg/router"
	"tasktracker-api/pkg/service"
)

func main() {
	repository := repository.NewRepository()
	services := service.NewService(repository)
	router := router.NewRouter(services)
	app := router.InitRoutes()
	app.Run(":8080")
}
