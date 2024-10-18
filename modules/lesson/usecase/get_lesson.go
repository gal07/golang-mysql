package usecase

import (
	"context"
	"gosql/modules/lesson/models"
	"gosql/modules/lesson/payload"
)

func (s *lessonUseCase) GetLesson(ctx context.Context, req payload.ReqGetAllLesson) (res payload.ResGetAllLesson, err error) {

	// call repository
	resp, err := s.lessonEntity.LessonRepo.Get(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

func (s *lessonUseCase) GetDetail(ctx context.Context, req payload.ReqGetDetail) (res models.Lesson, err error) {

	//call repository
	resp, err := s.lessonEntity.LessonRepo.GetDetail(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, err
}
