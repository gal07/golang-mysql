package usecase

import (
	"context"
	"gosql/modules/lesson/payload"
	teacherPayload "gosql/modules/teacher/payload"
	util "gosql/utility"
)

func (s *lessonUseCase) SearchLesson(ctx context.Context, req payload.ReqSearch) (res []payload.ResGetDetailLesson, err error) {

	// call repository
	resp, err := s.lessonEntity.LessonRepo.Search(ctx, req)
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
			res = append(res, ress)
		}
	}

	return res, err

}
