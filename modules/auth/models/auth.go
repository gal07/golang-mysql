package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	ID        int    `json:"id"`
	Username  string `json:"username" validate="required"`
	Email     string `json:"email" validate="required,email"`
	Password  string `json:"password" validate="required"`
	Salt      string `json:"salt"`
	Status    int    `json:"status"`
	Isdelete  int    `json:"isdelete"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
}

func (u *Auth) EncryptPassword() {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	u.Password = string(hashPassword)
}

func (u *Auth) ValidatePassword(password string) bool {
	fmt.Println("Upw : ", u.Password)
	fmt.Println("pw : ", password)

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false
	}
	return true
}
