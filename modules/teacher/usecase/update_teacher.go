package usecase

import (
	"context"
	"gosql/modules/teacher/models"
	"gosql/modules/teacher/payload"
)

func (s *teacherUseCase) UpdateTeacher(ctx context.Context, req models.Teacher) (res payload.ResUpdate, err error) {

	// Update
	if res, err = s.teacherEntity.TeacherRepo.Update(ctx, req); err != nil {
		return res, err
	}

	return res, err
}
