package database

import (
	"context"
	"database/sql"
	"gosql/modules/lesson"
	"gosql/modules/lesson/models"
	"gosql/modules/lesson/payload"

	_ "github.com/go-sql-driver/mysql"
)

type accessLesson struct {
	Db *sql.DB
}

func NewLessonRepo(Db *sql.DB) lesson.ILessonRepo {
	return accessLesson{Db: Db}
}

func (s accessLesson) Insert(ctx context.Context, req models.Lesson) (res models.Lesson, err error) {

	// Execute
	_, err = s.Db.Exec("insert into tb_lesson (name,teacher) values (?,?)", &req.Name, &req.Teacher)
	if err != nil {
		return res, err
	}
	return res, err
}

func (s accessLesson) Get(ctx context.Context, req payload.ReqGetAllLesson) (res []models.Lesson, err error) {
	// Instance
	var lesson []models.Lesson

	offset := (req.CurrentPage - 1) * req.PageSize
	// Execute
	rows, err := s.Db.Query("select id,name,teacher,status from tb_lesson WHERE isdelete = ? LIMIT ? OFFSET ?", 0, req.PageSize, offset)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var each = models.Lesson{}
		var err = rows.Scan(&each.ID, &each.Name, &each.Teacher, &each.Status)
		if err != nil {
			panic(err)
		} else {
			lesson = append(lesson, each)
		}
	}

	return lesson, err
}

func (s accessLesson) GetDetail(ctx context.Context, req payload.ReqGetDetail) (res models.Lesson, err error) {

	// Execute
	err = s.Db.QueryRow("select id,name,teacher,status from tb_lesson where id = ? and isdelete = ?", req.ID, 0).Scan(&res.ID, &res.Name, &res.Teacher, &res.Status)
	if err != nil {
		return res, err
	}

	return res, err
}

func (s accessLesson) Search(ctx context.Context, req payload.ReqSearch) (res []models.Lesson, err error) {

	// bind search value
	findteacher := "%" + req.Search + "%"

	// Querying to find spesific student
	rows, err := s.Db.Query("SELECT id,name,teacher,status from tb_lesson where name LIKE ? and isdelete = ?", findteacher, 0)
	for rows.Next() {
		var each = models.Lesson{}
		var err = rows.Scan(&each.ID, &each.Name, &each.Teacher, &each.Status)
		if err != nil {
			return res, err
		} else {
			res = append(res, each)
		}
	}
	return res, err
}

func (s accessLesson) Update(ctx context.Context, req models.Lesson) (res payload.ResUpdate, err error) {

	// Execute
	_, err = s.Db.Exec("UPDATE tb_lesson SET name = ?,teacher = ? WHERE id = ?", req.Name, req.Teacher, req.ID)
	if err != nil {
		return res, err
	}

	//Fill Struct
	res.Teacher = uint64(req.Teacher)
	res.Name = req.Name

	return res, err
}

func (s accessLesson) Delete(ctx context.Context, req payload.ReqDelete) (res bool, err error) {

	// Execute
	_, err = s.Db.Exec("UPDATE tb_lesson SET isdelete = ? WHERE id = ?", 1, req.ID)
	if err != nil {
		return false, err
	}

	return true, err
}
