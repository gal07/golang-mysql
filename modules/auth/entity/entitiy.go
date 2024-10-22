package entity

import (
	"database/sql"
	"gosql/modules/auth"
	"gosql/modules/auth/entity/database"

	_ "github.com/go-sql-driver/mysql"
)

type AuthEntity struct {
	AuthRepo auth.IAuthRepo
}

func NewEntity(Db *sql.DB) (AuthEntity, error) {
	return AuthEntity{
		AuthRepo: database.NewAuthRepo(Db)}, nil
}
