package usecase

import (
	"context"
	"gosql/modules/report/models"
	"gosql/modules/report/payload"
)

func (s *reportUseCase) UpdateReport(ctx context.Context, req models.Report) (res payload.ResUpdate, err error) {

	// Update
	if _, err := s.reportEntity.ReportRepo.Update(ctx, req); err != nil {
		return res, err
	}

	// Get name student, lesson, Teacher
	student, _, _ := getTheirNames(s, ctx, req.Student, req.Lesson)

	// Fill Response
	res.Student = student
	res.Grade = req.Grade
	res.ID = req.ID

	return res, nil
}
