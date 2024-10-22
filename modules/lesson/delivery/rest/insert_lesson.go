package rest

import (
	"gosql/modules/lesson/models"
	"gosql/modules/lesson/payload"
	util "gosql/utility"

	"github.com/gin-gonic/gin"
)

func (e endpoint) InsertLesson(c *gin.Context) {

	// Bind
	payloads := models.Lesson{}
	if err := c.Bind(&payloads); err != nil {
		panic(err)
	}

	// Validate Input Request
	notvalid := util.Validates(c, payloads)
	if notvalid {
		return
	}

	// Prevent duplicate teacher (one teacher only own one lesson)
	found, err := e.useCaseLesson.GetByTeacherID(c, payload.ReqByTeacherId{
		ID: payloads.Teacher,
	})
	if found {
		util.ResponseErrorCustom(c, 200, nil, "Teacher already owned lesson.")
		panic(err)
	}

	// call repository
	res, err := e.useCaseLesson.InsertLesson(c, payloads)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	util.ResponseOK(c, 200, res)
}
