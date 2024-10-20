package payload

type ReqGetAllReport struct {
	CurrentPage int `json:"current_page" validate:"required,numeric"`
	PageSize    int `json:"page_size" validate:"required,numeric"`
}

type ResGetAllReport struct {
	Current    int                  `json:"current_page"`
	NextPage   int                  `json:"next_page"`
	Pagesize   int                  `json:"page_size"`
	ListReport []ResGetDetailReport `json:"list_report"`
}

type ReqGetDetail struct {
	ID int
}

type ResGetDetailReport struct {
	ID        int    `json:"id"`
	Student   string `json:"student_name"`
	Lesson    string `json:"lesson_name"`
	Teacher   string `json:"teacher_name"`
	StudentID int    `json:"student_id"`
	LessonID  int    `json:"lesson_id"`
	Grade     int    `json:"grade"`
	Status    string `json:"status"`
}
