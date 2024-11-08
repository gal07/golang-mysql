package database

import (
	"context"
	"database/sql"
	"fmt"
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
	resp, err := s.Db.Exec("insert into tb_user (username,email,password,uuid) values(?,?,?,?)", &req.Username, &req.Email, &req.Password, &req.Uuid)
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

func (s accessAuth) InsertRegister(ctx context.Context, req models.Register) (res models.Register, err error) {
	// Execute
	resp, err := s.Db.Exec("insert into tb_register (uuid,username,email,password,activation_salt) values(?,?,?,?,?)", &req.Uuid, &req.Username, &req.Email, &req.Password, &req.ActivationSalt)
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

func (s accessAuth) FindRegisterBy(ctx context.Context, by string, targetby any) (res models.Register, err error) {

	// Execute
	err = s.Db.QueryRow("select id,uuid,username,email,password,activation_salt,is_activated from tb_register where "+by+" = ?", targetby).
		Scan(&res.ID, &res.Uuid, &res.Username, &res.Email, &res.Password, &res.ActivationSalt, &res.IsActivated)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s accessAuth) SetActived(ctx context.Context, req models.Register) (res models.Register, err error) {

	// Execute
	_, err = s.Db.Exec("UPDATE tb_register SET is_activated = ? WHERE activation_salt = ?", 1, req.ActivationSalt)
	if err != nil {
		return res, err
	}

	// Fill response
	res.Email = req.Email
	res.IsActivated = 1

	return res, nil

}

func (s accessAuth) DeleteRegister(ctx context.Context, req models.Register) (res bool, err error) {

	// Execute
	result, err := s.Db.Exec("DELETE from tb_register where id = ?", req.ID)
	if err != nil {
		return false, err
	}

	fmt.Println(result.RowsAffected())

	return true, nil

}
