package database

import (
	"context"
	"database/sql"
	"gosql/modules/teacher"
	"gosql/modules/teacher/models"
	"gosql/modules/teacher/payload"

	_ "github.com/go-sql-driver/mysql"
)

type accessTeachers struct {
	Db *sql.DB
}

func NewTeacherRepo(Db *sql.DB) teacher.ITeacherRepo {
	return accessTeachers{Db: Db}
}

func (s accessTeachers) Insert(ctx context.Context, req models.Teacher) (res models.Teacher, err error) {

	// Execute
	_, err = s.Db.Exec("insert into tb_teacher (name,age) values (?,?)", &req.Name, &req.Age)
	if err != nil {
		panic(err)
	}

	return req, err
}

func (s accessTeachers) Get(ctx context.Context, req payload.ReqGetAllTeacher) (res []models.Teacher, err error) {
	// Instance
	var teacher []models.Teacher
	offset := (req.CurrentPage - 1) * req.PageSize

	// Execute
	rows, err := s.Db.Query("select id,name,age,status,isdelete,created_at,deleted_at from tb_teacher WHERE isdelete = ? LIMIT ? OFFSET ?", 0, req.PageSize, offset)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var each = models.Teacher{}
		var err = rows.Scan(&each.ID, &each.Name, &each.Age, &each.Status, &each.Isdelete, &each.Createdat, &each.Deletedat)
		if err != nil {
			panic(err)
		} else {
			teacher = append(teacher, each)
		}
	}
	return teacher, err
}

func (s accessTeachers) GetDetail(ctx context.Context, req payload.ReqGetDetail) (res models.Teacher, err error) {

	err = s.Db.QueryRow("select id,name,age,status from tb_teacher where id = ? and isdelete = ?", req.ID, 0).Scan(&res.ID, &res.Name, &res.Age, &res.Status)
	if err != nil {
		return res, err
	}

	return res, err
}

func (s accessTeachers) Search(ctx context.Context, req payload.ReqSearch) (res []models.Teacher, err error) {

	// bind search value
	findteacher := "%" + req.Search + "%"

	// Querying to find spesific student
	rows, err := s.Db.Query("SELECT id,name,age,status from tb_teacher where name LIKE ? and isdelete = ?", findteacher, 0)
	for rows.Next() {
		var each = models.Teacher{}
		var err = rows.Scan(&each.ID, &each.Name, &each.Age, &each.Status)
		if err != nil {
			return res, err
		} else {
			res = append(res, each)
		}
	}
	return res, err
}

func (s accessTeachers) Update(ctx context.Context, req models.Teacher) (res payload.ResUpdate, err error) {

	// Execute
	_, err = s.Db.Exec("UPDATE tb_teacher SET name = ?,age = ? WHERE id = ?", req.Name, req.Age, req.ID)
	if err != nil {
		panic(err)
	}

	//Fill Struct
	res.Age = uint64(req.Age)
	res.Name = req.Name

	return res, err

}

func (s accessTeachers) Delete(ctx context.Context, req payload.ReqDelete) (res bool, err error) {

	// Execute
	_, err = s.Db.Exec("UPDATE tb_teacher SET isdelete = ? WHERE id = ?", 1, req.ID)
	if err != nil {
		return false, err
	}

	return true, err

}
