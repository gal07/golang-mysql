package usecase

import (
	"context"
	"gosql/modules/teacher/models"
	"gosql/modules/teacher/payload"
)

func (s *teacherUseCase) GetTeacher(ctx context.Context, req payload.ReqGetAllTeacher) (res payload.ResGetAllTeacher, err error) {

	// call repository
	resp, err := s.teacherEntity.TeacherRepo.Get(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

func (s *teacherUseCase) GetDetail(ctx context.Context, req payload.ReqGetDetail) (res models.Teacher, err error) {

	//call repository
	resp, err := s.teacherEntity.TeacherRepo.GetDetail(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, err
}
