package rest

import (
	"gosql/modules/report"

	"github.com/gin-gonic/gin"
)

type endpoint struct {
	useCaseReport report.IReportUseCase
}

func NewEndPoint(
	engine *gin.Engine,
	useCaseReport report.IReportUseCase,
) error {

	var edp = endpoint{
		useCaseReport: useCaseReport,
	}

	// Basic Auth
	const rootEndpoint = "/api/v1/report"
	r := engine.Group(rootEndpoint, gin.BasicAuth(gin.Accounts{
		"myuser": "pass123",
	}))

	r.POST("/get", edp.GetReport)
	r.POST("/create", edp.InsertReport)
	r.POST("/search", edp.SearchReport)
	r.PUT("/update/:id", edp.UpdateReport)
	r.DELETE("/delete/:id", edp.DeleteReport)
	r.GET("get/:id", edp.GetDetailReport)
	r.Use()

	return nil

}
