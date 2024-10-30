package rest

import (
	"gosql/modules/auth/payload"
	util "gosql/utility"

	"github.com/gin-gonic/gin"
)

func (e endpoint) RefreshToken(c *gin.Context) {

	// Bind
	payloads := payload.RefreshToken{}
	if err := c.BindJSON(&payloads); err != nil {
		panic(err)
	}

	// Validate Input Request
	notvalid := util.Validates(c, payloads)
	if notvalid {
		return
	}

	// call repository
	res, err := e.useCaseAuth.RefreshToken(c, payloads)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	util.ResponseOK(c, 200, res)
}
