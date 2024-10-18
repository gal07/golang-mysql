package models

import "time"

type JSONTime time.Time

type Teacher struct {
	ID        int    `json:"id"`
	Name      string `json:"name" validate:"required"`
	Age       int    `json:"age" validate:"required,gte=0,lte=50,numeric,number"`
	Status    int    `json:"status"`
	Isdelete  int    `json:"isdelete"`
	Createdat string `json:"created_at"`
	Deletedat string `json:"deleted_at"`
}
