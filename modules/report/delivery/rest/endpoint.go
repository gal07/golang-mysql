package rest

import (
	"gosql/modules/report"
	util "gosql/utility"

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
	const rootEndpoint = "/api/v1/reports"
	r := engine.Group(rootEndpoint, util.VerifyToken())
	r.GET("/", edp.GetReport)
	r.GET("/search", edp.SearchReport)
	r.GET("/:id", edp.GetDetailReport)
	r.POST("/", edp.InsertReport)
	r.PATCH("/:id", edp.UpdateReport)
	r.DELETE("/:id", edp.DeleteReport)
	r.Use()

	return nil

}
