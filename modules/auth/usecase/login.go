package usecase

import (
	"context"
	"errors"
	"gosql/modules/auth/payload"
	tokenModel "gosql/modules/token/models"
	util "gosql/utility"
)

func (s *authUseCase) Login(ctx context.Context, req payload.Login) (res payload.ResLogin, err error) {

	// Match email
	matchEmail, err := s.authEntity.AuthRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return res, err
	}

	// Validate password
	if !matchEmail.ValidatePassword(req.Password) {
		return res, errors.New("password not match")
	}

	// create token
	token, err := util.CreateToken(req.Email)
	if err != nil {
		return res, err
	}

	// create refresh token
	reftoken, expired, err := util.RefreshToken(req.Email)
	if err != nil {
		return res, err
	}

	// save refresh token to white list
	_, _ = s.tokenEntity.TokenRepo.Create(ctx, tokenModel.Token{
		Token:   reftoken,
		Type:    "refresh-token",
		Expires: expired,
		Email:   req.Email,
	})

	// save token to white lsit
	_, _ = s.tokenEntity.TokenRepo.Create(ctx, tokenModel.Token{
		Token:   token,
		Type:    "primary-token",
		Expires: expired,
		Email:   req.Email,
	})

	// fill response
	res.Token = token
	res.RefreshToken = reftoken
	res.Islogin = true

	return res, nil

}
