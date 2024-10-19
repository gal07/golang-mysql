package payload

type ResInsert struct {
	Name    string `json:"lesson_name"`
	Teacher string `json:"teacher_name"`
	Status  string `json:"status"`
}
