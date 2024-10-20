package payload

type ResInsert struct {
	Name   string `json:"teacher_name"`
	Age    int    `json:"teacher_age"`
	Status string `json:"status"`
}
