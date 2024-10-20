package usecase

import (
	"context"
	lessonPayload "gosql/modules/lesson/payload"
	"gosql/modules/report/payload"
	studentPayload "gosql/modules/students/payload"
	teacherPayload "gosql/modules/teacher/payload"
	util "gosql/utility"
)

func (s *reportUseCase) GetReport(ctx context.Context, req payload.ReqGetAllReport) (res payload.ResGetAllReport, err error) {

	// call repository
	resp, err := s.reportEntity.ReportRepo.Get(ctx, req)
	if err != nil {
		return res, err
	}

	// Fill Response
	for i := 0; i < len(resp); i++ {

		// Get Name Student Lesson & Teacher
		student, lesson, teacher := getTheirNames(s, ctx, resp[i].Student, resp[i].Lesson)

		// Fill response List detail report
		var ress = payload.ResGetDetailReport{
			ID:        resp[i].ID,
			Student:   student,
			Lesson:    lesson,
			Teacher:   teacher,
			StudentID: resp[i].Student,
			LessonID:  resp[i].Lesson,
			Grade:     resp[i].Grade,
			Status:    util.CheckStatus(resp[i].Status),
		}

		res.ListReport = append(res.ListReport, ress)

	}
	res.Current = req.CurrentPage
	res.NextPage = req.CurrentPage + 1
	res.Pagesize = req.PageSize

	return res, nil

}

func (s *reportUseCase) GetDetailReport(ctx context.Context, req payload.ReqGetDetail) (res payload.ResGetDetailReport, err error) {

	//call repository
	resp, err := s.reportEntity.ReportRepo.GetDetail(ctx, req)
	if err != nil {
		return res, err
	}

	// Get Name Student Lesson & Teacher
	student, lesson, teacher := getTheirNames(s, ctx, resp.Student, resp.Lesson)

	//Fill Response
	res.ID = resp.ID
	res.Lesson = lesson
	res.Status = util.CheckStatus(resp.Status)
	res.StudentID = resp.Student
	res.LessonID = resp.Lesson
	res.Student = student
	res.Teacher = teacher
	res.Grade = resp.Grade

	return res, err
}

func (s *reportUseCase) GetReportByStudentnLesson(ctx context.Context, req payload.ReqGetByStudentLesson) (res bool, err error) {
	// call repository
	get, err := s.reportEntity.ReportRepo.GetByStudentnLesson(ctx, req)
	if err != nil {
		return get, err
	}

	res = get

	return res, nil
}

// Get Name Student Lesson & Teacher
func getTheirNames(s *reportUseCase, ctx context.Context, StudentID int, LessonID int) (studentN string, lessonN string, teacherN string) {

	// Student
	var studentName = "-"
	student, err := s.studentEntity.StudentRepo.GetDetail(ctx, studentPayload.ReqGetDetail{ID: StudentID})
	if err != nil {
		studentName = "-"
	} else {
		studentName = student.Name
	}

	// Lesson
	var lessonName = "-"
	lesson, err := s.lessonEntity.LessonRepo.GetDetail(ctx, lessonPayload.ReqGetDetail{ID: LessonID})
	if err != nil {
		lessonName = "-"
	} else {
		lessonName = lesson.Name
	}

	// Teacher
	var teacherName = "-"
	teacher, err := s.teacherEntity.TeacherRepo.GetDetail(ctx, teacherPayload.ReqGetDetail{ID: lesson.Teacher})
	if err != nil {
		teacherName = "-"
	} else {
		teacherName = teacher.Name
	}

	return studentName, lessonName, teacherName
}
