package usecase

import (
	"context"
	"errors"
	"gosql/modules/report/models"
	"gosql/modules/report/payload"
)

func (s *reportUseCase) UpdateReport(ctx context.Context, req models.Report) (res payload.ResUpdate, err error) {

	// matching id ,student and lesson
	get, err := s.reportEntity.ReportRepo.GetDetail(ctx, payload.ReqGetDetail{ID: req.ID})
	if err != nil {
		return res, err
	}

	if get.Lesson != req.Lesson || get.Student != req.Student || get.ID != req.ID {
		return res, errors.New("request data is not match")
	}

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
