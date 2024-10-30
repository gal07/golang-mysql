package entity

import (
	"database/sql"
	"gosql/modules/token"
	"gosql/modules/token/entity/database"

	_ "github.com/go-sql-driver/mysql"
)

type TokenEntity struct {
	TokenRepo token.ITokenRepo
}

func NewEntity(Db *sql.DB) (TokenEntity, error) {
	return TokenEntity{
		TokenRepo: database.NewTokenRepo(Db)}, nil
}
