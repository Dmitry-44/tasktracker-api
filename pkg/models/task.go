package models

type Task struct {
	Id     int    `json:"id" db:"id" `
	Title  string `json:"title" db:"title"`
	Status int    `json:"status" db:"status"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

type TaskData struct {
	Title  *string `json:"title" db:"title"`
	Status *int    `json:"status" db:"status"`
}
