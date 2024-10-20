package usecase

import (
	"context"
	"gosql/modules/report/payload"
	util "gosql/utility"
)

func (s *reportUseCase) SearchReport(ctx context.Context, req payload.ReqSearch) (res []payload.ResGetDetailReport, err error) {

	// call repository
	resp, err := s.reportEntity.ReportRepo.Search(ctx, req)
	if err != nil {
		return res, err
	}

	// Fill Response
	for i := 0; i < len(resp); i++ {

		// Get Name Student Lesson & Teacher
		student, lesson, teacher := getTheirNames(s, ctx, resp[i].Student, resp[i].Lesson)

		// Fill response
		var ress = payload.ResGetDetailReport{
			ID:      resp[i].ID,
			Student: student,
			Lesson:  lesson,
			Teacher: teacher,
			Grade:   resp[i].Grade,
			Status:  util.CheckStatus(resp[i].Status),
		}

		res = append(res, ress)

	}

	return res, nil

}
