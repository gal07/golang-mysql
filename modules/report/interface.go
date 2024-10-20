package report

import (
	"context"
	"gosql/modules/report/models"
	"gosql/modules/report/payload"
)

type IReportUseCase interface {
	InsertReport(ctx context.Context, req models.Report) (res payload.ResInsert, err error)
	GetReport(ctx context.Context, req payload.ReqGetAllReport) (res payload.ResGetAllReport, err error)
	GetDetailReport(ctx context.Context, req payload.ReqGetDetail) (res payload.ResGetDetailReport, err error)
	GetReportByStudentnLesson(ctx context.Context, req payload.ReqGetByStudentLesson) (res bool, err error)
	SearchReport(ctx context.Context, req payload.ReqSearch) (res []payload.ResGetDetailReport, err error)
	UpdateReport(ctx context.Context, req models.Report) (res payload.ResUpdate, err error)
	DeleteReport(ctx context.Context, req payload.ReqDelete) (res bool, err error)
}

type IReportRepo interface {
	Insert(ctx context.Context, req models.Report) (res models.Report, err error)
	Get(ctx context.Context, req payload.ReqGetAllReport) (res []models.Report, err error)
	GetDetail(ctx context.Context, req payload.ReqGetDetail) (res models.Report, err error)
	GetByStudentnLesson(ctx context.Context, req payload.ReqGetByStudentLesson) (res bool, err error)
	Search(ctx context.Context, req payload.ReqSearch) (res []models.Report, err error)
	Update(ctx context.Context, req models.Report) (res models.Report, err error)
	Delete(ctx context.Context, req payload.ReqDelete) (res bool, err error)
}
