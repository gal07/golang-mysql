package usecase

import (
	"gosql/modules/lesson"
	"gosql/modules/lesson/entity"
	teacher "gosql/modules/teacher/entity"
)

type lessonUseCase struct {
	config        string
	lessonEntity  entity.LessonEntity
	teacherEntity teacher.TeacherEntity
}

func NewUseCase(config string, lessonEntity entity.LessonEntity, teacherEntity teacher.TeacherEntity) (lesson.ILessonUseCase, error) {
	return &lessonUseCase{
		config:        "test",
		lessonEntity:  lessonEntity,
		teacherEntity: teacherEntity}, nil
}
