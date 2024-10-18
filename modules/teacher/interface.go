package teacher

import (
	"context"
	"gosql/modules/teacher/models"
	"gosql/modules/teacher/payload"
)

type ITeacherUseCase interface {
	InsertTeacher(ctx context.Context, req models.Teacher) (res models.Teacher, err error)
	GetTeacher(ctx context.Context, req payload.ReqGetAllTeacher) (res payload.ResGetAllTeacher, err error)
	GetDetail(ctx context.Context, req payload.ReqGetDetail) (res models.Teacher, err error)
	SearchTeacher(ctx context.Context, req payload.ReqSearch) (res []models.Teacher, err error)
	UpdateTeacher(ctx context.Context, req models.Teacher) (res payload.ResUpdate, err error)
	DeleteTeacher(ctx context.Context, req payload.ReqDelete) (res bool, err error)
}

type ITeacherRepo interface {
	Insert(ctx context.Context, req models.Teacher) (res models.Teacher, err error)
	Get(ctx context.Context, req payload.ReqGetAllTeacher) (res payload.ResGetAllTeacher, err error)
	GetDetail(ctx context.Context, req payload.ReqGetDetail) (res models.Teacher, err error)
	Search(ctx context.Context, req payload.ReqSearch) (res []models.Teacher, err error)
	Update(ctx context.Context, req models.Teacher) (res payload.ResUpdate, err error)
	Delete(ctx context.Context, req payload.ReqDelete) (res bool, err error)
}
