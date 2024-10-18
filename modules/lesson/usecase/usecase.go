package usecase

import (
	"gosql/modules/lesson"
	"gosql/modules/lesson/entity"
)

type lessonUseCase struct {
	config       string
	lessonEntity entity.LessonEntity
}

func NewUseCase(config string, lessonEntity entity.LessonEntity) (lesson.ILessonUseCase, error) {
	return &lessonUseCase{
		config:       "test",
		lessonEntity: lessonEntity}, nil
}
