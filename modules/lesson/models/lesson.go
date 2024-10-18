package models

type Lesson struct {
	ID        int    `json:"id"`
	Name      string `json:"name" validate:"required"`
	Teacher   int    `json:"teacher" validate:"required,numeric,number"`
	Status    int    `json:"status"`
	Isdelete  int    `json:"isdelete"`
	Createdat string `json:"created_at"`
	Deletedat string `json:"deleted_at"`
}
