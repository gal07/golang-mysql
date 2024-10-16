package usecase

import (
	"context"
	"gosql/modules/students/models"
	"gosql/modules/students/payload"
)

func (s *studentUseCase) GetAllStudent(ctx context.Context, req payload.ReqGetAllStudents) (res payload.ResGetAllStudent, err error) {

	// call repository
	resp, err := s.studentEntity.StudentRepo.Get(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *studentUseCase) GetDetail(ctx context.Context, req payload.ReqGetDetail) (res models.Student, err error) {

	//call repository
	resp, err := s.studentEntity.StudentRepo.GetDetail(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, err
}
