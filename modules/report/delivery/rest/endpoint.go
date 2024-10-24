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
	r.Handle("POST", "/", edp.GetReport)
	r.POST("/create", edp.InsertReport)
	r.POST("/search", edp.SearchReport)
	r.PUT("/update/:id", edp.UpdateReport)
	r.DELETE("/delete/:id", edp.DeleteReport)
	r.GET("/:id", edp.GetDetailReport)
	r.Use()

	return nil

}
