package usecase

import (
	"context"
	"gosql/modules/students/models"
	"gosql/modules/students/payload"
)

func (s *studentUseCase) SearchStudent(ctx context.Context, req payload.ReqSearch) (res []models.Student, err error) {

	// call repository
	resp, err := s.studentEntity.StudentRepo.Search(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil

}
