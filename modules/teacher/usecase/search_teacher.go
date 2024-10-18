package usecase

import (
	"context"
	"gosql/modules/teacher/models"
	"gosql/modules/teacher/payload"
)

func (s *teacherUseCase) SearchTeacher(ctx context.Context, req payload.ReqSearch) (res []models.Teacher, err error) {

	// call repository
	resp, err := s.teacherEntity.TeacherRepo.Search(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil

}
