package usecase

import (
	"context"
	"gosql/modules/students/payload"
)

func (s *studentUseCase) MassInsertStudent(ctx context.Context, req payload.ReqMassCreate) (res payload.ResMassCreate, err error) {

	//insert
	if res, err := s.studentEntity.StudentRepo.MassInsert(ctx, req); err != nil {
		return res, err
	}
	return res, nil
}
