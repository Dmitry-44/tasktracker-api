package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"tasktracker-api/pkg/models"
)

type TaskRepo struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) GetAll(user int) (models.TaskList, error) {
	list := models.TaskList{}
	rows, err := r.db.Query("SELECT * FROM tasks WHERE created_by=($1)", user)
	if err != nil {
		return list, err
	}
	defer rows.Close()
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.Id, &task.Title, &task.Status, &task.CreatedBy, &task.Priority, &task.Description, &task.GroupId)
		if err != nil {
			log.Fatalf("scanning database error: %v", err)
		}
		list.Tasks = append(list.Tasks, task)
	}
	return list, nil
}

func (r *TaskRepo) GetTaskById(user int, taskId int) (models.Task, error) {
	task := models.Task{}
	err := r.db.QueryRow("SELECT * FROM tasks WHERE id=($1) created_by=($2)", taskId, user).Scan(&task.Id, &task.Title, &task.Status)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (r *TaskRepo) CreateTask(user int, task models.TaskData) (int, error) {
	var createdTaskId int
	set := make([]string, 0)
	numbersSet := make([]string, 0)
	values := make([]interface{}, 0)
	valueId := 1
	if task.Title != nil {
		set = append(set, "title")
		numbersSet = append(numbersSet, fmt.Sprintf("$%v", valueId))
		values = append(values, *task.Title)
		valueId++
	}
	if task.Description != nil {
		set = append(set, "description")
		numbersSet = append(numbersSet, fmt.Sprintf("$%v", valueId))
		values = append(values, *task.Description)
		valueId++
	}
	if task.Status != nil {
		set = append(set, "status")
		numbersSet = append(numbersSet, fmt.Sprintf("$%v", valueId))
		values = append(values, *task.Status)
		valueId++
	}
	if task.Priority != nil {
		set = append(set, "priority")
		numbersSet = append(numbersSet, fmt.Sprintf("$%v", valueId))
		values = append(values, *task.Priority)
		valueId++
	}
	if task.GroupId != nil {
		set = append(set, "group_id")
		numbersSet = append(numbersSet, fmt.Sprintf("$%v", valueId))
		values = append(values, *task.GroupId)
		valueId++
	}
	set = append(set, "created_by")
	numbersSet = append(numbersSet, fmt.Sprintf("$%v", valueId))
	values = append(values, user)
	setString := strings.Join(set, ", ")
	numbersSetString := strings.Join(numbersSet, ", ")
	query := fmt.Sprintf("INSERT into tasks (%s) VALUES (%s) RETURNING id", setString, numbersSetString)
	err := r.db.QueryRow(query, values...).Scan(&createdTaskId)
	if err != nil {
		return createdTaskId, err
	}
	return createdTaskId, nil
}

// TO DO
func (r *TaskRepo) UpdateTask(id int, task models.TaskData) error {
	set := make([]string, 0)
	args := make([]interface{}, 0)
	argsId := 1
	if task.Title != nil {
		set = append(set, fmt.Sprintf("title=($%d)", argsId))
		args = append(args, *task.Title)
		argsId++
	}
	if task.Status != nil {
		set = append(set, fmt.Sprintf("status=($%d)", argsId))
		args = append(args, *task.Status)
		argsId++
	}
	setQuery := strings.Join(set, ", ")
	query := fmt.Sprintf("UPDATE tasks SET %s WHERE id=($%v)", setQuery, argsId)
	args = append(args, id)
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskRepo) DeleteTask(id int) error {
	err := r.db.QueryRow("DELETE FROM tasks WHERE id=($1)", id)
	if err != nil {
		return err.Err()
	}
	return nil
}
