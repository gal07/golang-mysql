package entity

import (
	"database/sql"
	"gosql/modules/teacher"
	"gosql/modules/teacher/entity/database"

	_ "github.com/go-sql-driver/mysql"
)

type TeacherEntity struct {
	TeacherRepo teacher.ITeacherRepo
}

func NewEntity(Db *sql.DB) (TeacherEntity, error) {
	return TeacherEntity{
		TeacherRepo: database.NewTeacherRepo(Db)}, nil
}
