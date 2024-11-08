package auth

import (
	"context"
	"gosql/modules/auth/models"
	"gosql/modules/auth/payload"

	"github.com/gin-gonic/gin"
)

type IAuthUseCase interface {
	Login(ctx context.Context, req payload.Login) (res payload.ResLogin, err error)
	Register(ctx context.Context, req models.Register) (res payload.ResRegister, err error)
	RefreshToken(ctx *gin.Context, req payload.RefreshToken) (res payload.ResRefreshToken, err error)
	ActivatedUsers(ctx *gin.Context, req payload.ReqActivatedRegister) (res payload.ResActivatedRegister, err error)
}

type IAuthRepo interface {
	Insert(ctx context.Context, req models.Auth) (res models.Auth, err error)
	GetByEmail(ctx context.Context, email string) (res models.Auth, err error)
	GetByUsername(ctx context.Context, username string) (res models.Auth, err error)
	InsertRegister(ctx context.Context, req models.Register) (res models.Register, err error)
	FindRegisterBy(ctx context.Context, by string, targetby any) (res models.Register, err error)
	SetActived(ctx context.Context, req models.Register) (res models.Register, err error)
	DeleteRegister(ctx context.Context, req models.Register) (res bool, err error)
}
