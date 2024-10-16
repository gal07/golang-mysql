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
	engine.Group(rootEndpoint, gin.BasicAuth(gin.Accounts{
		"myuser": "pass123",
	}))

	engine.Handle("POST", rootEndpoint+"/get", edp.GetAllStudent)
	engine.Handle("POST", rootEndpoint+"/create", edp.InsertStudent)
	engine.Handle("POST", rootEndpoint+"/search", edp.SearchStudent)
	engine.Handle("POST", rootEndpoint+"/massinsert", edp.MassInsertStudent)
	engine.Handle("PUT", rootEndpoint+"/update/:id", edp.UpdateStudent)
	engine.Handle("DELETE", rootEndpoint+"/delete/:id", edp.DeleteStudent)
	engine.Handle("GET", rootEndpoint+"/get/:id", edp.GetStudent)
	engine.Use()
	return nil

}
