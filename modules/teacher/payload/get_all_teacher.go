package payload

type ReqGetAllTeacher struct {
	CurrentPage int `json:"current_page" validate:"required,numeric"`
	PageSize    int `json:"page_size" validate:"required,numeric"`
}

type ResGetAllTeacher struct {
	Current     int `json:"current_page"`
	NextPage    int `json:"next_page"`
	Pagesize    int `json:"page_size"`
	ListTeacher any `json:"list_teacher"`
}

type ReqGetDetail struct {
	ID int
}
