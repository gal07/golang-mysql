package rest

import (
	"gosql/modules/students"
	util "gosql/utility"

	"github.com/gin-gonic/gin"
)

type endpoint struct {
	useCaseStudent students.IStudentUseCase
}

func NewEndPoint(
	engine *gin.Engine,
	useCaseStudent students.IStudentUseCase,
) error {

	var edp = endpoint{
		useCaseStudent: useCaseStudent,
	}

	// Basic Auth
	const rootEndpoint = "/api/v1/students"
	r := engine.Group(rootEndpoint, util.VerifyToken())
	r.GET("/", edp.GetAllStudent)
	r.GET("/:id", edp.GetStudent)
	r.GET("/search", edp.SearchStudent)
	r.POST("/", edp.InsertStudent)
	r.POST("/massinsert", edp.MassInsertStudent)
	r.PATCH("/:id", edp.UpdateStudent)
	r.DELETE("/:id", edp.DeleteStudent)
	r.Use()

	return nil

}
