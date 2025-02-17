package rest

import (
	"gosql/modules/report/payload"
	util "gosql/utility"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (e endpoint) DeleteReport(c *gin.Context) {

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

	// call repository
	res, err := e.useCaseReport.DeleteReport(c, payload)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Failed")
		return
	}

	util.ResponseOK(c, 200, res)

}
