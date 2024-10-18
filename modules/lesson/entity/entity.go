package entity

import (
	"database/sql"
	"gosql/modules/lesson"
	"gosql/modules/lesson/entity/database"

	_ "github.com/go-sql-driver/mysql"
)

type LessonEntity struct {
	LessonRepo lesson.ILessonRepo
}

func NewEntity(Db *sql.DB) (LessonEntity, error) {
	return LessonEntity{
		LessonRepo: database.NewLessonRepo(Db)}, nil
}
