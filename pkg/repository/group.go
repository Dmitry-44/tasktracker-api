package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"tasktracker-api/pkg/models"
)

type GroupsRepo struct {
	db *sql.DB
}

func NewGroupsRepo(db *sql.DB) *GroupsRepo {
	return &GroupsRepo{db: db}
}

// func (r *TasksRepo) GetAll(user int) (models.TaskList, error) {
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

// func (r *TasksRepo) GetTaskById(user int, taskId int) (models.Task, error) {
// 	task := models.Task{}
// 	err := r.db.QueryRow("SELECT * FROM tasks WHERE id=($1) AND created_by=($2)", taskId, user).Scan(&task.Id, &task.Title, &task.Status, &task.CreatedBy, &task.Priority, &task.Description, &task.GroupId)
// 	if err != nil {
// 		return task, err
// 	}
// 	return task, nil
// }

func (r *GroupsRepo) CreateGroup(user int, group models.GroupData) (int, error) {
	var createdGroupId int
	set := make([]string, 0)
	numbersSet := make([]string, 0)
	values := make([]interface{}, 0)
	valueId := 1
	if group.Name != nil {
		set = append(set, "name")
		numbersSet = append(numbersSet, fmt.Sprintf("$%v", valueId))
		values = append(values, *group.Name)
		valueId++
	}
	if group.Description != nil {
		set = append(set, "description")
		numbersSet = append(numbersSet, fmt.Sprintf("$%v", valueId))
		values = append(values, *group.Description)
		valueId++
	}
	set = append(set, "created_by")
	numbersSet = append(numbersSet, fmt.Sprintf("$%v", valueId))
	values = append(values, user)
	setString := strings.Join(set, ", ")
	numbersSetString := strings.Join(numbersSet, ", ")
	tx, err := r.db.Begin()
	if err != nil {
		return createdGroupId, err
	}
	query := fmt.Sprintf("INSERT into groups (%s) VALUES (%s) RETURNING id", setString, numbersSetString)
	err = r.db.QueryRow(query, values...).Scan(&createdGroupId)
	if err != nil {
		return createdGroupId, err
	}
	err = r.SetUserGroup(user, createdGroupId)
	if err != nil {
		tx.Rollback()
		return createdGroupId, err
	}
	tx.Commit()
	return createdGroupId, nil
}

func (r *GroupsRepo) SetUserGroup(user int, group int) error {
	res, err := r.db.Exec("INSERT into public.user_group (user_id, group_id) VALUES ($1, $2)", user, group)
	if err != nil {
		return err
	}
	log.Printf("db: insert into public.user_group result: %v", res)
	return nil
}

// func (r *TasksRepo) UpdateTask(user int, id int, task models.TaskData) error {
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

// func (r *TasksRepo) DeleteTask(user int, id int) error {
// 	err := r.db.QueryRow("DELETE FROM tasks WHERE id=($1) and created_by=($2)", id, user)
// 	if err != nil {
// 		return err.Err()
// 	}
// 	return nil
// }
