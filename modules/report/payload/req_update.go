package payload

type ReqUpdate struct {
	ID    int `json:"id" validation:"required,numeric"`
	Grade int `json:"grade" validation:"required,numeric"`
}
