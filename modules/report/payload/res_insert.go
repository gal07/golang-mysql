package payload

type ResInsert struct {
	Student string `json:"student"`
	Lesson  string `json:"lesson"`
	Grade   int    `json:"grade"`
	Status  string `json:"status"`
}
