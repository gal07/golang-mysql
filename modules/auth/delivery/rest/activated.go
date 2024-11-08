package rest

import (
	"gosql/modules/auth/payload"
	util "gosql/utility"

	"github.com/gin-gonic/gin"
)

func (e endpoint) ActivatedUser(c *gin.Context) {

	// get query params
	activation := c.Query("activation")

	// Bind
	payload := payload.ReqActivatedRegister{
		ActivationSalt: activation,
	}

	// service
	notvalid := util.Validates(c, payload)
	if notvalid {
		return
	}

	// call repository
	res, err := e.useCaseAuth.ActivatedUsers(c, payload)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	util.ResponseOK(c, 200, res)

}
