package database

import (
	"database/sql"
	"fmt"
	"time"
)

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBname   string
	SSLmode  string
}

func NewDatabase(config DBConfig) (*sql.DB, error) {
	configString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.Username, config.Password, config.DBname, config.SSLmode)
	db, err := sql.Open("postgres", configString)
	if err != nil {
		return db, err
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(2 * time.Hour)
	return db, nil
}

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
