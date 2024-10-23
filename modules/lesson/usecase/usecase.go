package usecase

import (
	"gosql/modules/lesson"
	"gosql/modules/lesson/entity"
	report "gosql/modules/report/entity"
	teacher "gosql/modules/teacher/entity"
)

type lessonUseCase struct {
	config        string
	lessonEntity  entity.LessonEntity
	teacherEntity teacher.TeacherEntity
	reportEntity report.ReportEntity
}

func NewUseCase(config string, lessonEntity entity.LessonEntity, teacherEntity teacher.TeacherEntity, reportEntity report.ReportEntity) (lesson.ILessonUseCase, error) {
	return &lessonUseCase{
		config:        "test",
		lessonEntity:  lessonEntity,
		teacherEntity: teacherEntity,
		reportEntity: reportEntity,}, nil
}
