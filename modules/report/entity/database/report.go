package database

import (
	"context"
	"database/sql"
	"gosql/modules/report"
	"gosql/modules/report/models"
	"gosql/modules/report/payload"

	_ "github.com/go-sql-driver/mysql"
)

type accessReport struct {
	Db *sql.DB
}

func NewReportRepo(Db *sql.DB) report.IReportRepo {
	return accessReport{Db: Db}
}

func (s accessReport) Insert(ctx context.Context, req models.Report) (res models.Report, err error) {

	// Execute
	_, err = s.Db.Exec("insert into tb_report_card (student,lesson,grade) values(?,?,?)", &req.Student, &req.Lesson, &req.Grade)
	if err != nil {
		return res, err
	}

	return res, nil

}

func (s accessReport) Get(ctx context.Context, req payload.ReqGetAllReport) (res []models.Report, err error) {
	// Instance
	var report []models.Report
	offset := (req.CurrentPage - 1) * req.PageSize

	// Execute
	rows, err := s.Db.Query("select id,student,lesson,grade,status from tb_report_card WHERE isdelete = ? LIMIT ? OFFSET ?", 0, req.PageSize, offset)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var each = models.Report{}
		var err = rows.Scan(&each.ID, &each.Student, &each.Lesson, &each.Grade, &each.Status)
		if err != nil {
			panic(err)
		} else {
			report = append(report, each)
		}
	}

	return report, err
}

func (s accessReport) GetDetail(ctx context.Context, req payload.ReqGetDetail) (res models.Report, err error) {

	// Execute
	err = s.Db.QueryRow("select id,student,lesson,grade,status from tb_report_card where id = ? and isdelete = ?", req.ID, 0).Scan(&res.ID, &res.Student, &res.Lesson, &res.Grade, &res.Status)
	if err != nil {
		return res, err
	}

	return res, err
}

func (s accessReport) Search(ctx context.Context, req payload.ReqSearch) (res []models.Report, err error) {

	// bind search value
	findteacher := "%" + req.Search + "%"

	// Querying to find spesific student
	rows, err := s.Db.Query("SELECT id,student,lesson,grade,status from tb_report_card where grade LIKE ? and isdelete = ?", findteacher, 0)
	for rows.Next() {
		var each = models.Report{}
		var err = rows.Scan(&each.ID, &each.Student, &each.Lesson, &each.Grade, &each.Status)
		if err != nil {
			return res, err
		} else {
			res = append(res, each)
		}
	}
	return res, err
}

func (s accessReport) Update(ctx context.Context, req models.Report) (res models.Report, err error) {

	// Execute
	_, err = s.Db.Exec("UPDATE tb_report_card SET grade = ? WHERE id = ? AND student = ? AND lesson = ?", req.Grade, req.ID, req.Student, req.Lesson)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s accessReport) Delete(ctx context.Context, req payload.ReqDelete) (res bool, err error) {

	// Execute
	_, err = s.Db.Exec("UPDATE tb_report_card SET isdelete = ? WHERE id = ?", 1, req.ID)
	if err != nil {
		return false, err
	}

	return true, err
}

func (s accessReport) GetByStudentnLesson(ctx context.Context, req payload.ReqGetByStudentLesson) (res bool, err error) {
	// Execute
	var name = ""
	err = s.Db.QueryRow("select name from tb_report_card where student = ? lesson = ? and isdelete = ?", req.StudentID, req.LessonID, 0).Scan(&name)
	if err != nil {
		return true, nil
	}

	return false, nil
}
