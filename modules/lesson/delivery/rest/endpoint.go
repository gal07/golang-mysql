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
	const rootEndpoint = "/api/v1/lessons"
	r := engine.Group(rootEndpoint, util.VerifyToken())
	r.POST("/all", edp.GetLesson)
	r.POST("/", edp.InsertLesson)
	r.POST("/search", edp.SearchLesson)
	r.PUT("/:id", edp.UpdateLesson)
	r.DELETE("/:id", edp.DeleteLesson)
	r.GET("/:id", edp.GetDetail)
	r.Use()

	return nil

}
