package models

type Student struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name" validate:"required"`
	Age   uint64 `json:"age" validate:"required,gte=0,lte=50,numeric,number"`
	Grade uint64 `json:"grade" validate:"required,gte=0,lte=20,numeric,number"`
}
