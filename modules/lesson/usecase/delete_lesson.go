package usecase

import (
	"context"
	"gosql/modules/lesson/payload"
)

func (s *lessonUseCase) DeleteLesson(ctx context.Context, req payload.ReqDelete) (res bool, err error) {

	// Update
	if res, err = s.lessonEntity.LessonRepo.Delete(ctx, req); err != nil {
		return res, err
	}

	return res, err
}
