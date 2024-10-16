package usecase

import (
	"context"
	"gosql/modules/students/models"
	"gosql/modules/students/payload"
)

func (s *studentUseCase) InsertStudent(ctx context.Context, req payload.ReqInsert) (res payload.ReqInsert, err error) {

	// fill struct
	models := models.Student{
		Name:  req.Name,
		Age:   req.Age,
		Grade: req.Grade,
	}

	//insert
	if _, err := s.studentEntity.StudentRepo.Insert(ctx, models); err != nil {
		return req, err
	}

	return req, nil
}
