package main

import (
	"database/sql"
	"fmt"
	"log"
	"tasktracker-api/pkg/models"
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
	DBConfig := models.DBConfig{
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBname:   viper.GetString("db.dbname"),
		SSLmode:  viper.GetString("db.sslmode"),
	}
	// Opening a driver typically will not attempt to connect to the database.
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", DBConfig.Host, DBConfig.Port, DBConfig.Username, DBConfig.Password, DBConfig.DBname, DBConfig.SSLmode))
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//MIGRATIONS
	// m, err := migrate.New(
	// 	"file://db/migrations/",
	// 	"postgres://postgres:secret@localhost:5432/postgres?sslmode=disable")
	// if err != nil {
	// 	fmt.Println("migr err")
	// 	log.Fatal("migrate error: ", err)
	// }
	// if err := m.Up(); err != nil {
	// 	fmt.Println("migr err2")
	// 	log.Fatal(err)
	// }

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
