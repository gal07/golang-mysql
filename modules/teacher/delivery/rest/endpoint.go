package rest

import (
	"gosql/modules/teacher"

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
	const rootEndpoint = "/api/v1/teacher"
	r := engine.Group(rootEndpoint, gin.BasicAuth(gin.Accounts{
		"myuser": "pass123",
	}))

	r.POST("/get", edp.GetTeacher)
	r.POST("/create", edp.InsertTeacher)
	r.POST("/search", edp.SearchTeacher)
	r.PUT("/update/:id", edp.UpdateTeacher)
	r.DELETE("/delete/:id", edp.DeleteTeacher)
	r.GET("get/:id", edp.GetDetail)
	r.Use()

	return nil

}
