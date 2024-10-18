package usecase

import (
	"context"
	"gosql/modules/teacher/models"
)

func (s *teacherUseCase) InsertTeacher(ctx context.Context, req models.Teacher) (res models.Teacher, err error) {

	// insert
	if _, err := s.teacherEntity.TeacherRepo.Insert(ctx, req); err != nil {
		return req, err
	}

	return req, err
}
