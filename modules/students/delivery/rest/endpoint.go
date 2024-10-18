package rest

import (
	"gosql/modules/students"

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
	const rootEndpoint = "/api/v1/student"
	r := engine.Group(rootEndpoint, gin.BasicAuth(gin.Accounts{
		"myuser": "pass123",
	}))

	r.POST("/get", edp.GetAllStudent)
	r.POST("/create", edp.InsertStudent)
	r.POST("/search", edp.SearchStudent)
	r.PUT("/update/:id", edp.UpdateStudent)
	r.DELETE("/delete/:id", edp.DeleteStudent)
	r.GET("get/:id", edp.GetStudent)
	r.Use()

	return nil

}
