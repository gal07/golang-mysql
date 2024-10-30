package token

import (
	"context"
	"gosql/modules/token/models"
)

type ITokenUseCase interface {
}

type ITokenRepo interface {
	Create(ctx context.Context, req models.Token) (res models.Token, err error)
}
