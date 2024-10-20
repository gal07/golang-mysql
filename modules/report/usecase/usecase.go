package usecase

import (
	lesson "gosql/modules/lesson/entity"
	"gosql/modules/report"
	"gosql/modules/report/entity"
	students "gosql/modules/students/entity"
	teacher "gosql/modules/teacher/entity"
)

type reportUseCase struct {
	config        string
	reportEntity  entity.ReportEntity
	lessonEntity  lesson.LessonEntity
	teacherEntity teacher.TeacherEntity
	studentEntity students.StudentEntity
}

func NewUseCase(config string, reportEntity entity.ReportEntity, lessonEntity lesson.LessonEntity, teacherEntity teacher.TeacherEntity, studentEntity students.StudentEntity) (report.IReportUseCase, error) {
	return &reportUseCase{
		config:        "test",
		reportEntity:  reportEntity,
		lessonEntity:  lessonEntity,
		teacherEntity: teacherEntity,
		studentEntity: studentEntity}, nil
}
