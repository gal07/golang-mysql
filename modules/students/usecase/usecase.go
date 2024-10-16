package usecase

import (
	"gosql/modules/students"
	"gosql/modules/students/entity"
)

type studentUseCase struct {
	config        string
	studentEntity entity.StudentEntity
}

func NewUseCase(config string, studentEntity entity.StudentEntity) (students.IStudentUseCase, error) {
	return &studentUseCase{
		config:        "test",
		studentEntity: studentEntity}, nil
}
