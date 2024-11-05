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
	r.Handle("POST", "/all", edp.GetReport)
	r.POST("/", edp.InsertReport)
	r.POST("/search", edp.SearchReport)
	r.PUT("/:id", edp.UpdateReport)
	r.DELETE("/:id", edp.DeleteReport)
	r.GET("/:id", edp.GetDetailReport)
	r.Use()

	return nil

}
