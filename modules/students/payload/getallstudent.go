package payload

type ReqGetAllStudents struct {
	CurrentPage int `json:"current_page" validate:"required,numeric"`
	PageSize    int `json:"page_size" validate:"required,numeric"`
}

type ResGetAllStudent struct {
	Current     int `json:"current_page"`
	NextPage    int `json:"next_page"`
	Pagesize    int `json:"page_size"`
	ListStudent any `json:"list_student"`
}

type ReqGetDetail struct {
	ID int
}
