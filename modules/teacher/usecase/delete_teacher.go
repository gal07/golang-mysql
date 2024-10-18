package usecase

import (
	"context"
	"gosql/modules/teacher/payload"
)

func (s *teacherUseCase) DeleteTeacher(ctx context.Context, req payload.ReqDelete) (res bool, err error) {

	// Update
	if res, err = s.teacherEntity.TeacherRepo.Delete(ctx, req); err != nil {
		return res, err
	}

	return res, err
}
