package usecase

import (
	"context"
	"fmt"
	lessonPayload "gosql/modules/lesson/payload"
)

func (s *teacherUseCase) GetTeacherOnLesson(ctx context.Context, req lessonPayload.ReqByTeacherId) (res bool, err error) {

	// for prevent delete teacher if teacher id exist on tb_lesson
	get, err := s.lessonEntity.LessonRepo.GetByTeacherID(ctx, req)
	if err != nil {
		fmt.Println("test error : ",err)
		return false, err
	}
	_ = get

	return true, nil
}
