package rest

import (
	"gosql/modules/lesson/models"
	util "gosql/utility"

	"github.com/gin-gonic/gin"
)

func (e endpoint) InsertLesson(c *gin.Context) {

	// Bind
	payload := models.Lesson{}
	if err := c.Bind(&payload); err != nil {
		panic(err)
	}

	// Validate Input Request
	notvalid := util.Validates(c, payload)
	if notvalid {
		return
	}

	// call repository
	res, err := e.useCaseLesson.InsertLesson(c, payload)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	util.ResponseOK(c, 200, res)
}
