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
	payload := models.Lesson{}
	if err := c.Bind(&payload); err != nil {
		panic(err)
	}
	payload.ID = getID.ID

	// service
	notvalid := util.Validates(c, payload)
	if notvalid {
		return
	}

	// call repository
	res, err := e.useCaseLesson.UpdateLesson(c, payload)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	util.ResponseOK(c, 200, res)

}
