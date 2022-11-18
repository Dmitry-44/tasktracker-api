package models

type Group struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	CreatedBy   int    `json:"createdBy" db:"created_by"`
}

type GroupData struct {
	Name        *string `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
	CreatedBy   *int    `json:"createdBy" db:"created_by"`
}
