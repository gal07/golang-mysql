package models

import (
	"crypto/sha256"

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	ID        int    `json:"id"`
	Username  string `json:"username" validate="required,max=20,min=5"`
	Email     string `json:"email" validate="required,email"`
	Password  string `json:"password" validate="required"`
	Uuid      string `json:"uuid" validate="required"`
	Status    int    `json:"status"`
	Isdelete  int    `json:"isdelete"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
}

type Register struct {
	ID             int    `json:"id"`
	Uuid           string `json:"uuid"`
	Username       string `json:"username" validate="required,max=20,min=5"`
	Email          string `json:"email" validate="required,email"`
	Password       string `json:"password" validate="required"`
	ActivationSalt string `json:"activation_salt"`
	IsActivated    int    `json:"is_activated"`
	CreatedAt      int    `json:"created_at"`
	DeletedAt      int    `json:"deleted_at"`
}

func (u *Register) EncryptPasswords() {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	u.Password = string(hashPassword)

}

func (u *Auth) EncryptPassword() {
	// hashPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	// u.Password = string(hashPassword)
}

func (u *Register) EncryptActivationSalt(salt string) (res string) {
	Hash := sha256.New()
	Hash.Write([]byte(salt))
	return salt
}

func (u *Auth) ValidatePassword(password string) bool {

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false
	}
	return true

}
