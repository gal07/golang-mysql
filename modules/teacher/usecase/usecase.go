package usecase

import (
	lessonEntity "gosql/modules/lesson/entity"
	"gosql/modules/teacher"
	"gosql/modules/teacher/entity"
)

type teacherUseCase struct {
	config        string
	teacherEntity entity.TeacherEntity
	lessonEntity  lessonEntity.LessonEntity
}

func NewUseCase(config string, teacherEntity entity.TeacherEntity, lessonEntity lessonEntity.LessonEntity) (teacher.ITeacherUseCase, error) {
	return &teacherUseCase{
		config:        "test",
		teacherEntity: teacherEntity,
		lessonEntity:  lessonEntity,
	}, nil
}
