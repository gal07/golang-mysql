package usecase

import (
	"context"
	"errors"
	"fmt"
	"gosql/modules/auth/models"
	"gosql/modules/auth/payload"
	util "gosql/utility"
)

func (s *authUseCase) Login(ctx context.Context, req payload.Login) (res payload.ResLogin, err error) {

	// Match email
	matchEmail, err := s.authEntity.AuthRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return res, err
	}

	fmt.Println(matchEmail.Password)

	// Validate password
	if !matchEmail.ValidatePassword(req.Password) {
		return res, errors.New("password not match")
	}

	// create token
	token, err := util.CreateToken(req.Email)
	if err != nil {
		return res, err
	}

	// fill response
	res.Token = token
	res.Islogin = true

	return res, nil

}

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
