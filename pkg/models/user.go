package models

type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"-" binding:"required"`
	Groups   []int  `json:"groups" db:"groups"`
}

type UserData struct {
	Name     *string `json:"name" db:"name"`
	Username *string `json:"username" binding:"required"`
	Email    *string `json:"email" binding:"required"`
	Password *string `json:"password" binding:"required"`
	Groups   *[]int  `json:"groups" db:"groups"`
}

type AuthData struct {
	Username *string `json:"username" binding:"required"`
	Password *string `json:"password" binding:"required"`
}
