package token

import (
	"context"
	"gosql/modules/token/models"
	"gosql/modules/token/payload"
)

type ITokenUseCase interface {
}

type ITokenRepo interface {
	Create(ctx context.Context, req models.Token) (res models.Token, err error)
	Get(ctx context.Context, req payload.ResGetToken) (res models.Token, err error)
}
