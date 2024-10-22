package rest

import (
	"gosql/modules/lesson/models"
	"gosql/modules/lesson/payload"
	util "gosql/utility"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (e endpoint) UpdateLesson(c *gin.Context) {

	// Filtering URL
	id := c.Param("id")
	is, err := strconv.Atoi(id)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		return
	}

	// Fill Struct
	getID := payload.ReqGetDetail{
		ID: is,
	}

	// Bind
	payloads := models.Lesson{}
	if err := c.Bind(&payloads); err != nil {
		panic(err)
	}
	payloads.ID = getID.ID

	// service
	notvalid := util.Validates(c, payloads)
	if notvalid {
		return
	}

	// Prevent to update teacher ID where has been owned lesson
	found, err := e.useCaseLesson.GetByTeacherID(c, payload.ReqByTeacherId{
		ID: payloads.Teacher,
	})
	if found {
		util.ResponseErrorCustom(c, 200, nil, "Teacher already owned lesson.")
		panic(err)
	}

	// call repository
	res, err := e.useCaseLesson.UpdateLesson(c, payloads)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	util.ResponseOK(c, 200, res)

}
