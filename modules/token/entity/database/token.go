package database

import (
	"context"
	"database/sql"
	"gosql/modules/token"
	"gosql/modules/token/models"

	_ "github.com/go-sql-driver/mysql"
)

type accessToken struct {
	Db *sql.DB
}

func NewTokenRepo(Db *sql.DB) token.ITokenRepo {
	return accessToken{Db: Db}
}

// Create implements token.ITokenRepo.
func (s accessToken) Create(ctx context.Context, req models.Token) (res models.Token, err error) {

	// Execute
	_, err = s.Db.Exec("insert into tb_token_whitelist (token,type,expires) values(?,?,?)", req.Token, req.Type, req.Expires)
	if err != nil {
		return res, err
	}

	return res, nil

}

func (s accessToken) Delete(ctx context.Context, req models.Token) (res models.Token, err error) {

	// Execute
	_, err = s.Db.Exec("delete tb_token_whitelist where id = ?", req.ID)
	if err != nil {
		return res, err
	}
	return res, nil

}
