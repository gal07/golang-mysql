package rest

import (
	"gosql/modules/lesson/payload"
	util "gosql/utility"

	"github.com/gin-gonic/gin"
)

func (e endpoint) SearchLesson(c *gin.Context) {

	// get query params
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
	res, err := e.useCaseLesson.SearchLesson(c, payload)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	util.ResponseOK(c, 200, res)

}
