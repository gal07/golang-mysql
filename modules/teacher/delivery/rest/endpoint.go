package rest

import (
	"gosql/modules/teacher"
	util "gosql/utility"

	"github.com/gin-gonic/gin"
)

type endpoint struct {
	useCaseTeacher teacher.ITeacherUseCase
}

func NewEndPoint(
	engine *gin.Engine,
	useCaseTeacher teacher.ITeacherUseCase,
) error {

	var edp = endpoint{
		useCaseTeacher: useCaseTeacher,
	}

	// Basic Auth
	const rootEndpoint = "/api/v1/teachers"
	r := engine.Group(rootEndpoint, util.VerifyToken())
	r.POST("/all", edp.GetTeacher)
	r.POST("/", edp.InsertTeacher)
	r.POST("/search", edp.SearchTeacher)
	r.PUT("/:id", edp.UpdateTeacher)
	r.DELETE("/:id", edp.DeleteTeacher)
	r.GET("/:id", edp.GetDetail)
	r.Use()

	return nil

}
