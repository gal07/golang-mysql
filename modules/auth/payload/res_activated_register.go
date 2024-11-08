package payload

type ResActivatedRegister struct {
	Email       string `json:"email"`
	IsActivated string `json:"is_activated"`
}
