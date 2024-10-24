package rest

import (
	"gosql/modules/report/models"
	"gosql/modules/report/payload"
	util "gosql/utility"

	"github.com/gin-gonic/gin"
)

func (e endpoint) InsertReport(c *gin.Context) {

	// Bind
	payloads := models.Report{}
	if err := c.Bind(&payloads); err != nil {
		panic(err)
	}

	// Validate Input Request
	notvalid := util.Validates(c, payloads)
	if notvalid {
		return
	}

	// Check Every Student Have Grade Each Lesson if total of lesson is 5 then student must have report each lesson which is 5)
	found, err := e.useCaseReport.GetReportByStudentnLesson(c, payload.ReqGetByStudentLesson{StudentID: payloads.Student, LessonID: payloads.Lesson})
	if found {
		util.ResponseErrorCustom(c, 200, nil, "Can't entering same student & lesson")
		panic(err)
	}

	// call repository
	res, err := e.useCaseReport.InsertReport(c, payloads)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	util.ResponseOK(c, 200, res)
}
