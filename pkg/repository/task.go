package repository

import (
	"database/sql"
	"fmt"
	"log"
	"tasktracker-api/pkg/models"
)

type TaskRepo struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) GetAll(userId int) ([]models.Task, error) {
	var resp []models.Task = make([]models.Task, 0)
	// query := fmt.Sprintf("SELECT id, title FROM tasks")
	rows, err := r.db.Query("select * from tasks")
	// err := r.db.QueryRow(query).Scan(&resp)
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		// var id int
		// var title string
		err := rows.Scan(&task.Id, &task.Title)
		if err != nil {
			log.Fatalf("scanning database error: %v", err)
		}
		fmt.Printf("task in next is %v !!!", task)
		// fmt.Printf("title in next is %v", title)
		// resp.append(resp, task)
	}
	fmt.Printf("resp is, %v", resp)
	// fmt.Printf("res is, %s", res)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
// 	var id int
// 	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

// 	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
// 	if err := row.Scan(&id); err != nil {
// 		return 0, err
// 	}

// 	return id, nil
// }
