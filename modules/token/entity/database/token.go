package database

import (
	"context"
	"database/sql"
	"gosql/modules/token"
	"gosql/modules/token/models"
	"gosql/modules/token/payload"

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

func (s accessToken) Delete(ctx context.Context, req payload.DeleteToken) (res bool, err error) {

	// Execute
	_, err = s.Db.Exec("delete tb_token_whitelist where token = ?", req.Token)
	if err != nil {
		return false, err
	}
	return true, nil

}

func (s accessToken) Get(ctx context.Context, req payload.ResGetToken) (res models.Token, err error) {

	// Execute
	err = s.Db.QueryRow("select id,email,token,type,expires from tb_token_whitelist where token = ?", req.Token).Scan(&res.ID, &res.Email, &res.Token, &res.Type, &res.Expires)
	if err != nil {
		return res, err
	}

	return res, nil

}
