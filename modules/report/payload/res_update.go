package payload

type ResUpdate struct {
	ID      int    `json:"id"`
	Student string `json:"name"`
	Grade   int    `json:"grade"`
}
