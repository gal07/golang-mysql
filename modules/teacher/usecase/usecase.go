package usecase

import (
	"gosql/modules/teacher"
	"gosql/modules/teacher/entity"
)

type teacherUseCase struct {
	config        string
	teacherEntity entity.TeacherEntity
}

func NewUseCase(config string, teacherEntity entity.TeacherEntity) (teacher.ITeacherUseCase, error) {
	return &teacherUseCase{
		config:        "test",
		teacherEntity: teacherEntity}, nil
}
