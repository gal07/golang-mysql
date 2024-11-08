package payload

type ReqSetActivated struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	ActivationSalt string `json:"activation_salt"`
}
