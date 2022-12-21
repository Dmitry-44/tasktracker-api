package main

import (
	"fmt"
	"log"
	database "tasktracker-api/db"
	"tasktracker-api/pkg/repository"
	router "tasktracker-api/pkg/router"
	"tasktracker-api/pkg/service"

	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

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
	router := router.NewRouter(services)
	app := router.InitRoutes()
	// configCORS := cors.DefaultConfig()
	// configCORS.AllowCredentials = true
	// configCORS.AllowAllOrigins = true

	// configCORS.AllowOrigins = []string{"http://localhost:3001/"}

	// configCORS.AllowMethods = []string{"PUT", "PATCH", "GET", "POST"}
	// app.Use(cors.New(configCORS))

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:3001/"},
	// 	AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodHead, http.MethodDelete, http.MethodConnect},
	// 	AllowHeaders:     []string{"Accept", "Content-Type", "Origin", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control", "Access-Control-Allow-Origin"},
	// 	ExposeHeaders:    []string{"Content-Length", "Content-Type"},
	// 	AllowCredentials: true,
	// 	AllowWebSockets: true,
	// }))

	app.Run(viper.GetString("port"))
}

func getConfig() error {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
