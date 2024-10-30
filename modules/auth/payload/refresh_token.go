package payload

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
	Email        string `json:"email"`
}

type ResRefreshToken struct {
	Token  string `json:"token"`
	Email  string `json:"email"`
	Expire any    `json:"expired"`
}
