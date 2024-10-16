package entity

import (
	"database/sql"
	"gosql/modules/students"
	"gosql/modules/students/entity/database"

	_ "github.com/go-sql-driver/mysql"
)

type StudentEntity struct {
	StudentRepo students.IStudentRepo
}

func NewEntity(Db *sql.DB) (StudentEntity, error) {
	return StudentEntity{
		StudentRepo: database.NewStudentRepo(Db)}, nil
}
