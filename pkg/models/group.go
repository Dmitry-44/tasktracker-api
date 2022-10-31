package models

type Group struct {
	Id        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Users     []int  `json:"users" db:"users"`
	CreatedBy int    `json:"createdBy" db:"created_by"`
}
