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

func (r *TaskRepo) GetAll() (models.TaskList, error) {
	list := models.TaskList{}
	rows, err := r.db.Query("SELECT * FROM tasks")
	if err != nil {
		return list, err
	}
	defer rows.Close()
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.Id, &task.Title)
		if err != nil {
			log.Fatalf("scanning database error: %v", err)
		}
		list.Tasks = append(list.Tasks, task)
	}
	return list, nil
}

func (r *TaskRepo) GetTaskById(taskId int) (models.Task, error) {
	task := models.Task{}
	err := r.db.QueryRow("SELECT * FROM tasks WHERE id=($1)", taskId).Scan(&task.Id, &task.Title)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (r *TaskRepo) CreateTask(task models.TaskData) (int, error) {
	var taskId int
	query := "INSERT into tasks (title) VALUES ($1) RETURNING id"
	err := r.db.QueryRow(query, task.Title).Scan(&taskId)
	if err != nil {
		return taskId, err
	}
	return taskId, nil
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
