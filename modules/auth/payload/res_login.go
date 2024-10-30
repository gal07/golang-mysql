package payload

type ResLogin struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Islogin      bool   `json:"is_login"`
}
