package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"tasktracker-api/pkg/models"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

// func (r *UsersRepo) GetAll(user int) (models.TaskList, error) {
// 	taskList := models.TaskList{}
// 	taskList.Tasks = make([]models.Task, 0)
// 	rows, err := r.db.Query("SELECT * FROM tasks WHERE created_by=($1)", user)
// 	if err != nil {
// 		return taskList, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var task models.Task
// 		err := rows.Scan(&task.Id, &task.Title, &task.Status, &task.CreatedBy, &task.Priority, &task.Description, &task.GroupId)
// 		if err != nil {
// 			log.Fatalf("scanning database error: %v", err)
// 		}
// 		taskList.Tasks = append(taskList.Tasks, task)
// 	}
// 	return taskList, nil
// }

func (r *UsersRepo) GetUserById(id int) (models.User, error) {
	user := models.User{}
	var groupsArr string
	err := r.db.QueryRow("SELECT * FROM users WHERE id=($1)", id).Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &groupsArr)
	if err != nil {
		fmt.Printf("err db - %v", err.Error())
		return user, err
	}
	json.Unmarshal([]byte(groupsArr), &user.Groups)
	fmt.Printf("user from db - %v '\n'", user)
	return user, nil
}

func (r *UsersRepo) CreateUser(user models.UserData) (int, error) {
	var createdUserId int
	set := make([]string, 0)
	numbersSet := make([]string, 0)
	values := make([]interface{}, 0)
	valueId := 1
	if user.Name != nil {
		set = append(set, "name")
		numbersSet = append(numbersSet, fmt.Sprintf("$%v", valueId))
		values = append(values, *user.Name)
		valueId++
	}
	if user.Username != nil {
		set = append(set, "username")
		numbersSet = append(numbersSet, fmt.Sprintf("$%v", valueId))
		values = append(values, *user.Username)
		valueId++
	}
	if user.Password != nil {
		set = append(set, "password")
		numbersSet = append(numbersSet, fmt.Sprintf("$%v", valueId))
		values = append(values, *user.Password)
		valueId++
	}
	if user.Groups != nil {
		set = append(set, "groups")
		numbersSet = append(numbersSet, fmt.Sprintf("$%v", valueId))
		values = append(values, *user.Groups)
		valueId++
	}
	setString := strings.Join(set, ", ")
	numbersSetString := strings.Join(numbersSet, ", ")
	query := fmt.Sprintf("INSERT into users (%s) VALUES (%s) RETURNING id", setString, numbersSetString)
	err := r.db.QueryRow(query, values...).Scan(&createdUserId)
	if err != nil {
		return createdUserId, err
	}
	return createdUserId, nil
}

// func (r *UsersRepo) UpdateTask(user int, id int, task models.TaskData) error {
// 	set := make([]string, 0)
// 	args := make([]interface{}, 0)
// 	argsId := 1
// 	if task.Title != nil {
// 		set = append(set, fmt.Sprintf("title=($%d)", argsId))
// 		args = append(args, *task.Title)
// 		argsId++
// 	}
// 	if task.Status != nil {
// 		set = append(set, fmt.Sprintf("status=($%d)", argsId))
// 		args = append(args, *task.Status)
// 		argsId++
// 	}
// 	if task.Priority != nil {
// 		set = append(set, fmt.Sprintf("priority=($%d)", argsId))
// 		args = append(args, *task.Priority)
// 		argsId++
// 	}
// 	if task.Description != nil {
// 		set = append(set, fmt.Sprintf("description=($%d)", argsId))
// 		args = append(args, *task.Description)
// 		argsId++
// 	}
// 	setQuery := strings.Join(set, ", ")
// 	query := fmt.Sprintf("UPDATE tasks SET %s WHERE id=($%v) AND created_by=($%v)", setQuery, argsId, argsId+1)
// 	args = append(args, id)
// 	args = append(args, user)
// 	_, err := r.db.Exec(query, args...)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (r *UsersRepo) DeleteTask(user int, id int) error {
// 	err := r.db.QueryRow("DELETE FROM tasks WHERE id=($1) and created_by=($2)", id, user)
// 	if err != nil {
// 		return err.Err()
// 	}
// 	return nil
// }
