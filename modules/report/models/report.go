package models

type Report struct {
	ID      int `json:"id"`
	Student int `json:"student" validate:"required"`
	Lesson  int `json:"lesson" validate:"required"`
	Grade   int `json:"grade" validate:"required"`
	Status  int `json:"status"`
}
