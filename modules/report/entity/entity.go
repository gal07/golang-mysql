package entity

import (
	"database/sql"
	"gosql/modules/report"
	"gosql/modules/report/entity/database"

	_ "github.com/go-sql-driver/mysql"
)

type ReportEntity struct {
	ReportRepo report.IReportRepo
}

func NewEntity(Db *sql.DB) (ReportEntity, error) {
	return ReportEntity{
		ReportRepo: database.NewReportRepo(Db)}, nil
}
