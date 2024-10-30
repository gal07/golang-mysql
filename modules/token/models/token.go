package models

type Token struct {
	ID      uint64 `json:"id"`
	Email   string `json:"email" validate:"required,email"`
	Token   string `json:"token" validate:"required"`
	Expires any    `json:"expires" validate:"required,numeric,number"`
	Type    string `json:"type" validate:"required,alpha"`
}
