package rest

import (
	"gosql/modules/lesson"
	util "gosql/utility"

	"github.com/gin-gonic/gin"
)

type endpoint struct {
	useCaseLesson lesson.ILessonUseCase
}

func NewEndPoint(
	engine *gin.Engine,
	useCaseLesson lesson.ILessonUseCase,
) error {

	var edp = endpoint{
		useCaseLesson: useCaseLesson,
	}

	// Basic Auth
	const rootEndpoint = "/api/v1/lesson"
	r := engine.Group(rootEndpoint, util.VerifyToken())
	r.POST("/get", edp.GetLesson)
	r.POST("/create", edp.InsertLesson)
	r.POST("/search", edp.SearchLesson)
	r.PUT("/update/:id", edp.UpdateLesson)
	r.DELETE("/delete/:id", edp.DeleteLesson)
	r.GET("get/:id", edp.GetDetail)
	r.Use()

	return nil

}
