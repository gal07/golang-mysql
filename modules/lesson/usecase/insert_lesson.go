package usecase

import (
	"context"
	"gosql/modules/lesson/models"
)

func (s *lessonUseCase) InsertLesson(ctx context.Context, req models.Lesson) (res models.Lesson, err error) {

	// insert
	if _, err := s.lessonEntity.LessonRepo.Insert(ctx, req); err != nil {
		return req, err
	}

	return req, err
}
