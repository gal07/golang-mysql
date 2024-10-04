package models

type Responses struct {
	Error   error  `json:"error"`
	Message string `json:"message"`
}
