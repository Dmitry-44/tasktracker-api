package models

type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Username string `json:"username" binding:"required"`
	Password string `json:"-" binding:"required"`
	Groups   []int  `json:"groups" db:"groups"`
}
