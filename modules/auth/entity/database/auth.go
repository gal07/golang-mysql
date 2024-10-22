package database

import (
	"context"
	"database/sql"
	"gosql/modules/auth"
	"gosql/modules/auth/models"

	_ "github.com/go-sql-driver/mysql"
)

type accessAuth struct {
	Db *sql.DB
}

func NewAuthRepo(Db *sql.DB) auth.IAuthRepo {
	return accessAuth{Db: Db}
}

func (s accessAuth) Insert(ctx context.Context, req models.Auth) (res models.Auth, err error) {

	// Execute
	resp, err := s.Db.Exec("insert into tb_user (username,email,password,salt) values(?,?,?,?)", &req.Username, &req.Email, &req.Password, &req.Salt)
	if err != nil {
		return res, err
	}

	// Fill Response
	id, err := resp.LastInsertId()
	if err != nil {

	}

	res.ID = int(id)

	return res, nil

}

func (s accessAuth) GetByEmail(ctx context.Context, email string) (res models.Auth, err error) {

	// Execute
	err = s.Db.QueryRow("select id,username,email,status,password from tb_user where email = ?", email).Scan(&res.ID, &res.Username, &res.Email, &res.Status, &res.Password)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s accessAuth) GetByUsername(ctx context.Context, username string) (res models.Auth, err error) {

	// Execute
	err = s.Db.QueryRow("select id,username,email,status from tb_user where username = ?", username).Scan(&res.ID, &res.Username, &res.Email, &res.Status)
	if err != nil {
		return res, err
	}

	return res, nil
}
