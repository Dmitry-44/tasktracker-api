package models

type Task struct {
	Id    int    `json:"id" db:"id" `
	Title string `json:"title" db:"title"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

type TaskData struct {
	Title *string `json:"title" db:"title"`
}
