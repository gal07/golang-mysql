package usecase

import (
	"context"
	"gosql/modules/lesson/payload"
	teacherPayload "gosql/modules/teacher/payload"
	util "gosql/utility"
)

func (s *lessonUseCase) GetLesson(ctx context.Context, req payload.ReqGetAllLesson) (res payload.ResGetAllLesson, err error) {

	// call repository
	resp, err := s.lessonEntity.LessonRepo.Get(ctx, req)
	if err != nil {
		return res, err
	}

	// Fill Response
	for i := 0; i < len(resp); i++ {
		name, err := s.teacherEntity.TeacherRepo.GetDetail(ctx, teacherPayload.ReqGetDetail{ID: resp[i].Teacher})
		if err != nil {
			return res, err
		} else {
			var ress = payload.ResGetDetailLesson{
				ID:      resp[i].ID,
				Lesson:  resp[i].Name,
				Teacher: name.Name,
				Status:  util.CheckStatus(resp[i].Status),
			}
			res.Listlesson = append(res.Listlesson, ress)
		}
	}
	res.Current = req.CurrentPage
	res.NextPage = req.CurrentPage + 1
	res.Pagesize = req.PageSize

	return res, nil

}

func (s *lessonUseCase) GetDetail(ctx context.Context, req payload.ReqGetDetail) (res payload.ResGetDetailLesson, err error) {

	//call repository
	resp, err := s.lessonEntity.LessonRepo.GetDetail(ctx, req)
	if err != nil {
		return res, err
	}

	//Fill Response
	teacher, err := s.teacherEntity.TeacherRepo.GetDetail(ctx, teacherPayload.ReqGetDetail{ID: resp.Teacher})
	if err != nil {
		return res, err
	}
	res.ID = resp.ID
	res.Lesson = resp.Name
	res.Teacher = teacher.Name
	res.Status = util.CheckStatus(resp.Status)

	return res, err
}

func (s *lessonUseCase) GetByTeacherID(ctx context.Context, req payload.ReqByTeacherId) (res bool, err error) {
	// call repository
	_, err = s.lessonEntity.LessonRepo.GetByTeacherID(ctx, req)
	if err != nil {
		return false, err
	}

	return true, err
}
