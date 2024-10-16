package usecase

import (
	"context"
	"gosql/modules/students/payload"
)

func (s *studentUseCase) DeleteStudent(ctx context.Context, req payload.ReqDelete) (res bool, err error) {

	// Update
	if res, err = s.studentEntity.StudentRepo.Delete(ctx, req); err != nil {
		return res, err
	}

	return res, err
}
