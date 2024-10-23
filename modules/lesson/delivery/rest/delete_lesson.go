package rest

import (
	"gosql/modules/lesson/payload"
	reportPayload "gosql/modules/report/payload"
	util "gosql/utility"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (e endpoint) DeleteLesson(c *gin.Context) {

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

	// Prevent delete lesson if id lesson exist in tb_report_card
	isfound, err := e.useCaseLesson.GetByLesson(c, reportPayload.GetByLesson{
		LessonID: payload.ID,
	})

	if isfound {
		util.ResponseErrorCustom(c, 200, nil, "Can't delete lesson, lesson exist on active record.")
		return
	}

	// call repository
	res, err := e.useCaseLesson.DeleteLesson(c, payload)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Failed")
		return
	}

	util.ResponseOK(c, 200, res)

}
