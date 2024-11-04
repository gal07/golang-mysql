package usecase

import (
	"errors"
	"gosql/modules/auth/payload"
	tokenModels "gosql/modules/token/models"
	util "gosql/utility"

	"github.com/gin-gonic/gin"
)

func (s *authUseCase) RefreshToken(ctx *gin.Context, req payload.RefreshToken) (res payload.ResRefreshToken, err error) {

	// check valid token
	isvalid := util.ValidToken(ctx, req.RefreshToken)
	if !isvalid {
		util.CreateLog(errors.New("Token not valid."))
		return res, errors.New("Token not valid.")
	}

	// do refresh token
	refreshToken, expires, err := util.RefreshToken(req.Email)
	if err != nil {
		return res, err
	}

	// Fill response
	res.Email = req.Email
	res.Token = refreshToken
	res.Expire = expires

	// save token
	_, err = s.tokenEntity.TokenRepo.Create(ctx, tokenModels.Token{
		Email:   res.Email,
		Token:   res.Token,
		Expires: res.Expire,
		Type:    "primary-token",
	})

	return res, nil

}
