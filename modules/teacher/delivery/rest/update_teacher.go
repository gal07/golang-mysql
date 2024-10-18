package rest

import (
	"gosql/modules/teacher/models"
	"gosql/modules/teacher/payload"
	util "gosql/utility"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (e endpoint) UpdateTeacher(c *gin.Context) {

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
	payload := models.Teacher{}
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
	res, err := e.useCaseTeacher.UpdateTeacher(c, payload)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	util.ResponseOK(c, 200, res)

}
