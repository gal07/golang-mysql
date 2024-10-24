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
	r.POST("/", edp.GetAllStudent)
	r.POST("/create", edp.InsertStudent)
	r.POST("/massinsert", edp.MassInsertStudent)
	r.POST("/search", edp.SearchStudent)
	r.PUT("/update/:id", edp.UpdateStudent)
	r.DELETE("/delete/:id", edp.DeleteStudent)
	r.GET("/:id", edp.GetStudent)
	r.Use()

	return nil

}
