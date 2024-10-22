package payload

type ResLogin struct {
	Token   string `json:"token"`
	Islogin bool   `json:"is_login"`
}
