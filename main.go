package main

import (
	"database/sql"
	"log"
	"tasktracker-api/pkg/repository"
	router "tasktracker-api/pkg/router"
	"tasktracker-api/pkg/service"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

func main() {
	err := getConfig()
	if err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	// Opening a driver typically will not attempt to connect to the database.
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=secret dbname=postgres sslmode=disable")
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	defer db.Close()
	//MIGRATIONS
	m, err := migrate.New(
		"/db/migrations",
		"postgres://postgres:secret@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("migrate error: ", err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	router := router.NewRouter(services)
	app := router.InitRoutes()
	app.Run(viper.GetString("port"))
}

func getConfig() error {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
