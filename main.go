package main

import (
	"fmt"
	"log"
	database "tasktracker-api/db"
	"tasktracker-api/pkg/hub"
	"tasktracker-api/pkg/repository"
	router "tasktracker-api/pkg/router"
	"tasktracker-api/pkg/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

type userCtx string

const ctxKeyUser userCtx = "user"

func main() {
	err := getConfig()
	if err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	DBConfig := database.DBConfig{
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBname:   viper.GetString("db.dbname"),
		SSLmode:  viper.GetString("db.sslmode"),
	}
	// Opening a driver typically will not attempt to connect to the database.
	db, err := database.NewDatabase(DBConfig)
	if err != nil {
		log.Fatal(fmt.Printf("init database error : %v", err))
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(fmt.Printf("ping database error : %v", err))
	}
	defer db.Close()

	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	hub := hub.NewHub()
	go hub.Run()
	router := router.NewRouter(services, hub)
	app := router.InitRoutes()
	app.GET("/ws", func(c *gin.Context) {
		router.WSHandler(hub, c.Writer, c.Request)
	})
	app.Run(viper.GetString("port"))
}

func getConfig() error {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
