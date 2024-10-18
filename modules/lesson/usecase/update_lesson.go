package usecase

import (
	"context"
	"gosql/modules/lesson/models"
	"gosql/modules/lesson/payload"
)

func (s *lessonUseCase) UpdateLesson(ctx context.Context, req models.Lesson) (res payload.ResUpdate, err error) {

	// Update
	if res, err = s.lessonEntity.LessonRepo.Update(ctx, req); err != nil {
		return res, err
	}

	return res, err
}
