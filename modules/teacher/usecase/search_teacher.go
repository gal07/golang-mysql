package usecase

import (
	"context"
	lessonPayload "gosql/modules/lesson/payload"
	"gosql/modules/teacher/payload"
	util "gosql/utility"
)

func (s *teacherUseCase) SearchTeacher(ctx context.Context, req payload.ReqSearch) (res []payload.ResGetDetailTeacher, err error) {

	// call repository
	resp, err := s.teacherEntity.TeacherRepo.Search(ctx, req)
	if err != nil {
		return res, err
	}

	// Fill response
	for i := 0; i < len(resp); i++ {
		var lessonName = ""
		lesson, err := s.lessonEntity.LessonRepo.GetByTeacherID(ctx, lessonPayload.ReqByTeacherId{ID: resp[i].ID})
		if err != nil {
			lessonName = "-"
		} else {
			lessonName = lesson.Name
		}

		var ress = payload.ResGetDetailTeacher{
			ID:     resp[i].ID,
			Lesson: lessonName,
			Name:   resp[i].Name,
			Status: util.CheckStatus(resp[i].Status),
			Age:    resp[i].Age,
		}
		res = append(res, ress)
	}

	return res, nil

}
