package rest

import (
	"gosql/modules/auth/models"
	util "gosql/utility"

	"github.com/gin-gonic/gin"
)

func (e endpoint) Register(c *gin.Context) {

	// Bind
	payloads := models.Auth{}
	if err := c.Bind(&payloads); err != nil {
		panic(err)
	}

	// Validate Input Request
	notvalid := util.Validates(c, payloads)
	if notvalid {
		return
	}

	// call repository
	res, err := e.useCaseAuth.Register(c, payloads)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	util.ResponseOK(c, 200, res)
}
