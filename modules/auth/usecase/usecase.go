package usecase

import (
	"gosql/modules/auth"
	"gosql/modules/auth/entity"
)

type authUseCase struct {
	config     string
	authEntity entity.AuthEntity
}

func NewUseCase(config string, authEntity entity.AuthEntity) (auth.IAuthUseCase, error) {
	return &authUseCase{
		config:     "test",
		authEntity: authEntity}, nil
}
