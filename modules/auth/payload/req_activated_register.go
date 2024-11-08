package payload

type ReqActivatedRegister struct {
	Email          string `json:email`
	ActivationSalt string `json:activation_salt validate="required"`
}
