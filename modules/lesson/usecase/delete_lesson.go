package usecase

import (
	"context"
	"gosql/modules/lesson/payload"
	reportPayload "gosql/modules/report/payload"
)

func (s *lessonUseCase) DeleteLesson(ctx context.Context, req payload.ReqDelete) (res bool, err error) {

	// Update
	if res, err = s.lessonEntity.LessonRepo.Delete(ctx, req); err != nil {
		return res, err
	}

	return res, err
}

// Prevent delete lesson if, lesson id exist on tb_report_card
func (s *lessonUseCase) GetByLesson(ctx context.Context, req reportPayload.GetByLesson) (res bool, err error) {
	// call repository
	get, err := s.reportEntity.ReportRepo.GetByLesson(ctx, req)
	if err != nil {
		return false, err
	}

	if get.ID == 0 {
		return false, err
	}

	return true, nil
}
