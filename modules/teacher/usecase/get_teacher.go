package usecase

import (
	"context"
	lessonPayload "gosql/modules/lesson/payload"
	"gosql/modules/teacher/payload"
	util "gosql/utility"
)

func (s *teacherUseCase) GetTeacher(ctx context.Context, req payload.ReqGetAllTeacher) (res payload.ResGetAllTeacher, err error) {

	// call repository
	resp, err := s.teacherEntity.TeacherRepo.Get(ctx, req)
	if err != nil {
		return res, err
	}

	// Fill Response
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
		res.ListTeacher = append(res.ListTeacher, ress)
	}

	res.Current = req.CurrentPage
	res.NextPage = req.CurrentPage + 1
	res.Pagesize = req.PageSize

	return res, nil

}

func (s *teacherUseCase) GetDetail(ctx context.Context, req payload.ReqGetDetail) (res payload.ResGetDetailTeacher, err error) {

	//call repository
	resp, err := s.teacherEntity.TeacherRepo.GetDetail(ctx, req)
	if err != nil {
		return res, err
	}

	// Fill Response
	var lessonName = ""
	lesson, err := s.lessonEntity.LessonRepo.GetByTeacherID(ctx, lessonPayload.ReqByTeacherId{ID: resp.ID})
	if err != nil {
		lessonName = "-"
	} else {
		lessonName = lesson.Name
	}

	res.Name = resp.Name
	res.Age = resp.Age
	res.Lesson = lessonName
	res.ID = resp.ID
	res.Status = util.CheckStatus(resp.Status)

	return res, nil
}
