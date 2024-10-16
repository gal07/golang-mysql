package rest

import (
	"gosql/modules/students/payload"
	util "gosql/utility"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (e endpoint) GetAllStudent(c *gin.Context) {

	// Bind
	payload := payload.ReqGetAllStudents{}
	if err := c.Bind(&payload); err != nil {
		panic(err)
	}

	// service
	notvalid := util.Validates(c, payload)
	if notvalid {
		return
	}

	// call repository
	res, err := e.useCaseStudent.GetAllStudent(c, payload)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	util.ResponseOK(c, 200, res)

}

func (e endpoint) GetStudent(c *gin.Context) {

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
	res, err := e.useCaseStudent.GetDetail(c, payload)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	util.ResponseOK(c, 200, res)

}
