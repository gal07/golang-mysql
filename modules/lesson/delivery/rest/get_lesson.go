package rest

import (
	"gosql/modules/lesson/payload"
	util "gosql/utility"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (e endpoint) GetLesson(c *gin.Context) {

	// Bind
	payload := payload.ReqGetAllLesson{}
	if err := c.Bind(&payload); err != nil {
		panic(err)
	}

	// service
	notvalid := util.Validates(c, payload)
	if notvalid {
		return
	}

	// call repository
	res, err := e.useCaseLesson.GetLesson(c, payload)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	util.ResponseOK(c, 200, res)

}

func (e endpoint) GetDetail(c *gin.Context) {

	// Filtering URL
	id := c.Param("id")
	is, err := strconv.Atoi(id)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
	}

	// Fill Struct
	payload := payload.ReqGetDetail{
		ID: is,
	}

	// call repository
	res, err := e.useCaseLesson.GetDetail(c, payload)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	util.ResponseOK(c, 200, res)

}
