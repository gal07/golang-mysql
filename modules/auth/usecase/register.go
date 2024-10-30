package usecase

import (
	"context"
	"gosql/modules/auth/models"
)

func (s *authUseCase) Register(ctx context.Context, req models.Auth) (res models.Auth, err error) {

	// Encrypt password
	req.EncryptPassword()

	// Insert
	register, err := s.authEntity.AuthRepo.Insert(ctx, req)
	if err != nil {
		return res, err
	}

	res = register

	return res, nil

}
