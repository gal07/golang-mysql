package usecase

import (
	"gosql/modules/auth"
	"gosql/modules/auth/entity"
	tokenEntity "gosql/modules/token/entity"
)

type authUseCase struct {
	config      string
	authEntity  entity.AuthEntity
	tokenEntity tokenEntity.TokenEntity
}

func NewUseCase(config string, authEntity entity.AuthEntity, tokenEntity tokenEntity.TokenEntity) (auth.IAuthUseCase, error) {
	return &authUseCase{
		config:      "test",
		authEntity:  authEntity,
		tokenEntity: tokenEntity}, nil
}
