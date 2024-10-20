package usecase

import (
	"context"
	"gosql/modules/report/models"
	"gosql/modules/report/payload"
	util "gosql/utility"
)

func (s *reportUseCase) InsertReport(ctx context.Context, req models.Report) (res payload.ResInsert, err error) {

	// insert
	if _, err := s.reportEntity.ReportRepo.Insert(ctx, req); err != nil {
		return res, err
	}

	// Get Name Student Lesson & Teacher
	student, lesson, _ := getTheirNames(s, ctx, req.Student, req.Lesson)

	// Fill Response
	res.Student = student
	res.Lesson = lesson
	res.Status = util.CheckStatus(1)
	res.Grade = req.Grade

	return res, err
}
