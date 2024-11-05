package rest

import (
	"gosql/modules/report/payload"
	util "gosql/utility"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (e endpoint) GetReport(c *gin.Context) {

	// get query params
	cp, err := strconv.Atoi(c.Query("current_page"))
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}
	ps, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	// Bind
	payload := payload.ReqGetAllReport{
		CurrentPage: cp,
		PageSize:    ps,
	}

	// service
	notvalid := util.Validates(c, payload)
	if notvalid {
		return
	}

	// call repository
	res, err := e.useCaseReport.GetReport(c, payload)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	util.ResponseOK(c, 200, res)

}

func (e endpoint) GetDetailReport(c *gin.Context) {

	// Filtering URL
	id := c.Param("id")
	is, err := strconv.Atoi(id)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	// Fill Struct
	payload := payload.ReqGetDetail{
		ID: is,
	}

	// call repository
	res, err := e.useCaseReport.GetDetailReport(c, payload)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	util.ResponseOK(c, 200, res)

}
