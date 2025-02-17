package database

import (
	"context"
	"database/sql"
	"gosql/modules/students"
	"gosql/modules/students/models"
	"gosql/modules/students/payload"
	util "gosql/utility"

	_ "github.com/go-sql-driver/mysql"
)

type accessStudents struct {
	Db *sql.DB
}

func NewStudentRepo(Db *sql.DB) students.IStudentRepo {
	return accessStudents{Db: Db}
}

func (s accessStudents) Insert(ctx context.Context, req models.Student) (res models.Student, err error) {

	// Execute
	_, err = s.Db.Exec("insert into tb_student (name,age,grade) values (?, ?, ?)", &req.Name, &req.Age, &req.Grade)
	if err != nil {
		panic(err)
	}
	return req, err

}

func (s accessStudents) Get(ctx context.Context, req payload.ReqGetAllStudents) (res payload.ResGetAllStudent, err error) {
	// Instance
	var student []models.Student

	offset := (req.CurrentPage - 1) * req.PageSize
	// Execute
	rows, err := s.Db.Query("select id,name,age,grade from tb_student WHERE isdelete = ? LIMIT ? OFFSET ?", 0, req.PageSize, offset)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var each = models.Student{}
		var err = rows.Scan(&each.ID, &each.Name, &each.Age, &each.Grade)
		if err != nil {
			panic(err)
		} else {
			student = append(student, each)
		}
	}

	// Fill Response
	res.ListStudent = student
	res.Current = req.CurrentPage
	res.NextPage = req.CurrentPage + 1
	res.Pagesize = req.PageSize

	return res, err

}

func (s accessStudents) GetDetail(ctx context.Context, req payload.ReqGetDetail) (res models.Student, err error) {

	err = s.Db.QueryRow("select id,name,age,grade from tb_student where id = ? and isdelete = ?", req.ID, 0).Scan(&res.ID, &res.Name, &res.Age, &res.Grade)
	if err != nil {
		return res, err
	}

	return res, err
}

func (s accessStudents) Search(ctx context.Context, req payload.ReqSearch) (res []models.Student, err error) {
	// bind search value
	findstudent := "%" + req.Search + "%"

	// Querying to find spesific student
	rows, err := s.Db.Query("SELECT id,name,grade,age from tb_student where name LIKE ? and isdelete = ?", findstudent, 0)
	for rows.Next() {
		var each = models.Student{}
		var err = rows.Scan(&each.ID, &each.Name, &each.Grade, &each.Age)
		if err != nil {
			return res, err
		} else {
			res = append(res, each)
		}
	}
	return res, err
}

func (s accessStudents) Update(ctx context.Context, req models.Student) (res payload.ResUpdate, err error) {

	// Execute
	_, err = s.Db.Exec("UPDATE tb_student SET name = ?,age = ?,grade = ? WHERE id = ?", req.Name, req.Age, req.Grade, req.ID)
	if err != nil {
		panic(err)
	}

	//Fill Struct
	res.Age = req.Age
	res.Name = req.Name
	res.Grade = req.Grade

	return res, err

}

func (s accessStudents) Delete(ctx context.Context, req payload.ReqDelete) (res bool, err error) {

	// Execute
	_, err = s.Db.Exec("UPDATE tb_student SET isdelete = ? WHERE id = ?", 1, req.ID)
	if err != nil {
		return false, err
	}

	return true, err

}

func (s accessStudents) MassInsert(ctx context.Context, req payload.ReqMassCreate) (res payload.ResMassCreate, err error) {

	// Transaction begin
	tx, err := s.Db.BeginTx(ctx, nil)
	if err != nil {
		return res, err
	}
	var markinsert = 0
	for i := 0; i < req.Total; i++ {
		_, err := tx.ExecContext(ctx, "insert into tb_student (name,age,grade) values (?, ?, ?)", util.RandomFullname(), util.RandomID(20), util.RandomID(20))
		if err != nil {
			return res, err
		} else {
			markinsert = i + 1
		}
	}

	err = tx.Commit()
	if err != nil {
		return res, err
	}

	res.TotalInput = markinsert

	return res, err
}
