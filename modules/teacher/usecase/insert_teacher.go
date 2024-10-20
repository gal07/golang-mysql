package usecase

import (
	"context"
	"gosql/modules/teacher/models"
	"gosql/modules/teacher/payload"
	util "gosql/utility"
)

func (s *teacherUseCase) InsertTeacher(ctx context.Context, req models.Teacher) (res payload.ResInsert, err error) {

	// insert
	if _, err := s.teacherEntity.TeacherRepo.Insert(ctx, req); err != nil {
		return res, err
	}

	// Fill Response
	res.Name = req.Name
	res.Age = req.Age
	res.Status = util.CheckStatus(1)

	return res, nil
}
