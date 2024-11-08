package usecase

import (
	"errors"
	"gosql/modules/auth/models"
	"gosql/modules/auth/payload"
	util "gosql/utility"

	"github.com/gin-gonic/gin"
)

func (s *authUseCase) ActivatedUsers(ctx *gin.Context, req payload.ReqActivatedRegister) (res payload.ResActivatedRegister, err error) {

	// Make sure salt is valid
	claims, valid := util.ValidToken(ctx, req.ActivationSalt)
	if !valid {
		return res, errors.New("activation not valid, try register again")
	}

	// get claims value from token
	md := claims.(map[string]string)

	// Check email if exist tb_user
	_, err = s.authEntity.AuthRepo.GetByEmail(ctx, md["email"])
	if err == nil {
		return res, errors.New("email already active")
	}

	// Make sure email exist in tb_register
	matchemail, err := s.authEntity.AuthRepo.FindRegisterBy(ctx, "email", md["email"])
	if err != nil {
		return res, errors.New("email not exist")
	}

	// Set Active user
	resp, err := s.authEntity.AuthRepo.SetActived(ctx, models.Register{
		Email:          md["email"],
		ActivationSalt: matchemail.ActivationSalt,
		Username:       md["username"],
	})
	if err != nil {
		return res, err
	}

	// move user from tb_register to tb_user and delete data in tb_register
	_, err = s.authEntity.AuthRepo.Insert(ctx, models.Auth{
		Username: matchemail.Username,
		Email:    matchemail.Email,
		Password: matchemail.Password,
		Uuid:     matchemail.Uuid,
	})

	// Delete user in tb_register
	delete, err := s.authEntity.AuthRepo.DeleteRegister(ctx, models.Register{
		ID: matchemail.ID,
	})
	if !delete {
		return res, errors.New("delete failed")
	}

	// Fill Response
	res.Email = resp.Email
	res.IsActivated = "Activated"

	return res, nil

}
