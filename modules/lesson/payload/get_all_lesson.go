package payload

type ReqGetAllLesson struct {
	CurrentPage int `json:"current_page" validate:"required,numeric"`
	PageSize    int `json:"page_size" validate:"required,numeric"`
}

type ResGetAllLesson struct {
	Current    int `json:"current_page"`
	NextPage   int `json:"next_page"`
	Pagesize   int `json:"page_size"`
	Listlesson any `json:"list_lesson"`
}

type ReqGetDetail struct {
	ID int
}
