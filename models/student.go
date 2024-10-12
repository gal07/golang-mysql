package models

type Student struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name" validate:"required"`
	Age   uint64 `json:"age" validate:"required,gte=0,lte=50,numeric,number"`
	Grade uint64 `json:"grade" validate:"required,gte=0,lte=20,numeric,number"`
}

type Search struct {
	Search string `json:"search" validate:"required,alpha"`
}

type MassInsert struct {
	Total int `json:"total" validate:"required,numeric,number,gte=0,lte=500000"`
}
type Delete struct {
	ID int `json:"id"`
}
