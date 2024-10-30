package rest

import (
	"gosql/modules/students/payload"
	util "gosql/utility"

	"github.com/gin-gonic/gin"
)

func (e endpoint) MassInsertStudent(c *gin.Context) {

	// Bind
	payload := payload.ReqMassCreate{}
	if err := c.Bind(&payload); err != nil {
		panic(err)
	}

	// Validate Input Request
	notvalid := util.Validates(c, payload)
	if notvalid {
		return
	}

	// call repository
	res, err := e.useCaseStudent.MassInsertStudent(c, payload)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	util.ResponseOK(c, 200, res)
}
