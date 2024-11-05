package payload

type ResRegister struct {
	Email         string `json:"email"`
	Username      string `json:"username"`
	Is_registered bool   `json:"is_registered"`
	Is_active     bool   `json:"is_active"`
}
