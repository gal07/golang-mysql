package rest

import (
	"gosql/modules/teacher/payload"
	util "gosql/utility"

	"github.com/gin-gonic/gin"
)

func (e endpoint) SearchTeacher(c *gin.Context) {

	// Query params
	search := c.Query("search")

	// Bind
	payload := payload.ReqSearch{
		Search: search,
	}

	// service
	notvalid := util.Validates(c, payload)
	if notvalid {
		return
	}

	// call repository
	res, err := e.useCaseTeacher.SearchTeacher(c, payload)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	util.ResponseOK(c, 200, res)

}
