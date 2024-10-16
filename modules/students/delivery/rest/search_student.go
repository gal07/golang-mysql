package rest

import (
	"gosql/modules/students/payload"
	util "gosql/utility"

	"github.com/gin-gonic/gin"
)

func (e endpoint) SearchStudent(c *gin.Context) {

	// Bind
	payload := payload.ReqSearch{}
	if err := c.Bind(&payload); err != nil {
		panic(err)
	}

	// service
	notvalid := util.Validates(c, payload)
	if notvalid {
		return
	}

	// call repository
	res, err := e.useCaseStudent.SearchStudent(c, payload)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	util.ResponseOK(c, 200, res)

}
