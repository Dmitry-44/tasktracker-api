package models

type Task struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
}
