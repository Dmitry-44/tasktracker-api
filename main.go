package main

import (
	"database/sql"
	"log"
	"tasktracker-api/pkg/repository"
	router "tasktracker-api/pkg/router"
	"tasktracker-api/pkg/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"

	"github.com/spf13/viper"
)

func main() {
	//congig init
	err := getConfig()
	if err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	//db connection
	db, err := sql.Open("mysql", getDbConfigString())
	if err != nil {
		log.Fatalf("error db connection: %s", err.Error())
	}
	defer db.Close()

	//migrations
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"mysql",
		driver,
	)
	m.Up()

	repository := repository.NewRepository()
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

func getDbConfigString() string {
	return viper.GetString("db.username") + ":" + viper.GetString("db.password") + "@tcp(" + viper.GetString("db.host") + ":" + viper.GetString("db.port") + ")/" + viper.GetString("db.dbname")
}
