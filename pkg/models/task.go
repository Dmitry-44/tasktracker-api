package models

type Task struct {
	Id          int    `json:"id" db:"id" `
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Status      int    `json:"status" db:"status"`
	Priority    int    `json:"priority" db:"priority"`
	CreatedBy   int    `json:"createdBy" db:"created_by"`
	GroupId     int    `json:"groupId" db:"group_id"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

type TaskData struct {
	Title       *string `json:"title" db:"title"`
	Description *string `json:"description" db:"description"`
	Status      *int    `json:"status" db:"status"`
	Priority    *int    `json:"priority" db:"priority"`
	GroupId     *int    `json:"groupId" db:"group_id"`
}
