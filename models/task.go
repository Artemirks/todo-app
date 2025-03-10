package models

type Task struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Title     string `json:"title" validate:"required"`
	Completed bool   `json:"completed"`
}
