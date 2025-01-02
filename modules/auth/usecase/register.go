package usecase

import (
	"context"
	"errors"
	"gosql/modules/auth/models"
	"gosql/modules/auth/payload"
	util "gosql/utility"
	"runtime"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s *authUseCase) Register(ctx context.Context, req models.Register) (res payload.ResRegister, err error) {

	runtime.GOMAXPROCS(runtime.NumCPU())

	// Check email if exist tb_user
	_, err = s.authEntity.AuthRepo.GetByEmail(ctx, req.Email)
	if err == nil {
		return res, errors.New("email already exist")
	}

	// check email if exist in tb_register
	_, err = s.authEntity.AuthRepo.FindRegisterBy(ctx, "email", req.Email)
	if err == nil {
		return res, errors.New("email already exist")
	}

	// Prepare Request
	req.Uuid = util.RandomUUID()
	req.EncryptPasswords()

	// Prepare activation salt token
	activation, err := util.CreateToken(jwt.MapClaims{
		"username": req.Username,
		"email":    req.Email,
		"exp":      time.Now().Add(time.Minute * 1).Unix(),
		"type":     "activate-token",
	})
	if err != nil {
		return res, err
	}
	req.ActivationSalt = activation

	// Insert
	_, err = s.authEntity.AuthRepo.InsertRegister(ctx, req)
	if err != nil {
		return res, err
	}

	// send email
	var mails = []string{req.Email}
	var cc = []string{req.Email}
	go util.Send(mails, cc, activation)

	// Fill payload
	res.Email = req.Email
	res.Username = req.Username
	res.Is_active = false
	res.Is_registered = true

	return res, nil

}
