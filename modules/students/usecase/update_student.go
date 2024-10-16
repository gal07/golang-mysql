package usecase

import (
	"context"
	"gosql/modules/students/models"
	"gosql/modules/students/payload"
)

func (s *studentUseCase) UpdateStudent(ctx context.Context, req models.Student) (res payload.ResUpdate, err error) {

	// Update
	if res, err = s.studentEntity.StudentRepo.Update(ctx, req); err != nil {
		return res, err
	}

	return res, err
}
