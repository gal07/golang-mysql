package usecase

import (
	"context"
	"gosql/modules/auth/models"
	"gosql/modules/auth/payload"
	util "gosql/utility"
	"runtime"
)

func (s *authUseCase) Register(ctx context.Context, req models.Auth) (res payload.ResRegister, err error) {

	runtime.GOMAXPROCS(runtime.NumCPU())

	// Encrypt password
	req.EncryptPassword()

	// Insert
	_, err = s.authEntity.AuthRepo.Insert(ctx, req)
	if err != nil {
		return res, err
	}

	// send email
	var mails = []string{}
	var cc = []string{}
	mails = append(mails, req.Email)
	cc = append(cc, req.Email)
	go util.Send(mails, cc)

	// Fill payload
	res.Email = req.Email
	res.Username = req.Username
	res.Is_active = false
	res.Is_registered = true

	return res, nil

}
