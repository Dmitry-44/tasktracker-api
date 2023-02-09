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

func (r *GroupsRepo) GetAll(user int) (models.GroupList, error) {
	groupList := models.GroupList{}
	groupList.Groups = make([]models.Group, 0)
	rows, err := r.db.Query(
		`SELECT public."groups".id, public."groups".name, public."groups".description, public."groups".created_by 
		FROM public."groups" 
		LEFT JOIN user_group ON public."groups".id=user_group.group_id 
		WHERE user_id=($1);`, user)
	if err != nil {
		return groupList, err
	}
	defer rows.Close()
	for rows.Next() {
		var group models.Group
		err := rows.Scan(&group.Id, &group.Name, &group.Description, &group.CreatedBy)
		if err != nil {
			log.Fatalf("scanning database error: %v", err)
		}
		groupList.Groups = append(groupList.Groups, group)
	}
	return groupList, nil
}

func (r *GroupsRepo) GetGroupById(id int) (models.Group, error) {
	group := models.Group{}
	err := r.db.QueryRow("SELECT * FROM groups WHERE id=($1)", id).Scan(&group.Id, &group.Name, &group.Description, &group.CreatedBy)
	if err != nil {
		return group, err
	}
	return group, nil
}

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
	err = tx.QueryRow(query, values...).Scan(&createdGroupId)
	if err != nil {
		_ = tx.Rollback()
		return createdGroupId, err
	}
	res, err := tx.Exec("INSERT into public.user_group (user_id, group_id) VALUES ($1, $2) ON CONFLICT DO NOTHING", user, createdGroupId)
	_ = res
	if err != nil {
		_ = tx.Rollback()
		return createdGroupId, err
	}
	if err = tx.Commit(); err != nil {
		return createdGroupId, err
	}
	return createdGroupId, nil
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

func (r *GroupsRepo) DeleteGroupById(id int) error {
	res, err := r.db.Exec(`DELETE FROM public."groups" WHERE id=($1);`, id)
	_ = res
	if err != nil {
		return err
	}
	return nil
}

func (r *GroupsRepo) IsUserInGroup(userId int, groupId int) (bool, error) {
	var userFromDb int
	var groupFromDb int
	res := false
	err := r.db.QueryRow("SELECT * FROM user_group WHERE user_id=($1) and group_id=($2)", userId, groupId).Scan(&userFromDb, &groupFromDb)
	if err != nil {
		return res, err
	}
	if userFromDb == userId {
		return true, nil
	}
	return res, nil
}
