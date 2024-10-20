package payload

type ReqDelete struct {
	ID int `json:"id" validate:"required,numeric"`
}
