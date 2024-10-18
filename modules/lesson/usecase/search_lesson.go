package usecase

import (
	"context"
	"gosql/modules/lesson/models"
	"gosql/modules/lesson/payload"
)

func (s *lessonUseCase) SearchLesson(ctx context.Context, req payload.ReqSearch) (res []models.Lesson, err error) {

	// call repository
	resp, err := s.lessonEntity.LessonRepo.Search(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil

}
