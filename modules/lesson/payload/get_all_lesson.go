package payload

type ReqGetAllLesson struct {
	CurrentPage int `json:"current_page" validate:"required,numeric"`
	PageSize    int `json:"page_size" validate:"required,numeric"`
}

type ResGetAllLesson struct {
	Current    int                  `json:"current_page"`
	NextPage   int                  `json:"next_page"`
	Pagesize   int                  `json:"page_size"`
	Listlesson []ResGetDetailLesson `json:"list_lesson"`
}

type ReqGetDetail struct {
	ID int
}

type ResGetDetailLesson struct {
	ID      int    `json:"id"`
	Lesson  string `json:"lesson_name"`
	Teacher string `json:"teacher_name"`
	Status  string `json:"status"`
}
