package rest

import (
	"gosql/modules/teacher/payload"
	util "gosql/utility"
	"strconv"

	lessonPayload "gosql/modules/lesson/payload"

	"github.com/gin-gonic/gin"
)

func (e endpoint) DeleteTeacher(c *gin.Context) {

	// Filtering URL
	id := c.Param("id")
	is, err := strconv.Atoi(id)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		return
	}

	// Bind
	payload := payload.ReqDelete{}
	if err := c.Bind(&payload); err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		return
	}

	// Match
	if is != payload.ID {
		util.ResponseErrorCustom(c, 200, nil, "ID Not Match")
		return
	}

	// service
	notvalid := util.Validates(c, payload)
	if notvalid {
		return
	}

	// Prevent delete teacher, if teacher id found on tb_lesson
	isfound, err := e.useCaseTeacher.GetTeacherOnLesson(c, lessonPayload.ReqByTeacherId{
		ID: payload.ID,
	})
	if isfound {
		util.ResponseErrorCustom(c, 200, nil, "Can't delete teacher, teacher exist on active record.")
		return
	}

	// call repository
	res, err := e.useCaseTeacher.DeleteTeacher(c, payload)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Failed")
		return
	}

	util.ResponseOK(c, 200, res)

}
