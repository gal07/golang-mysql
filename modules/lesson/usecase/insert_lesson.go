package usecase

import (
	"context"
	"gosql/modules/lesson/models"
	"gosql/modules/lesson/payload"
	teacherPayload "gosql/modules/teacher/payload"
	util "gosql/utility"
)

func (s *lessonUseCase) InsertLesson(ctx context.Context, req models.Lesson) (res payload.ResInsert, err error) {

	// insert
	if _, err := s.lessonEntity.LessonRepo.Insert(ctx, req); err != nil {
		return res, err
	}

	// Fill Response
	name, err := s.teacherEntity.TeacherRepo.GetDetail(ctx, teacherPayload.ReqGetDetail{ID: req.Teacher})
	if err != nil {
		return res, err
	}
	res.Name = req.Name
	res.Teacher = name.Name
	res.Status = util.CheckStatus(1)

	return res, err
}
