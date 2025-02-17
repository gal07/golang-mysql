package teacher

import (
	"context"
	lessonPayload "gosql/modules/lesson/payload"
	"gosql/modules/teacher/models"
	"gosql/modules/teacher/payload"
)

type ITeacherUseCase interface {
	InsertTeacher(ctx context.Context, req models.Teacher) (res payload.ResInsert, err error)
	GetTeacher(ctx context.Context, req payload.ReqGetAllTeacher) (res payload.ResGetAllTeacher, err error)
	GetDetail(ctx context.Context, req payload.ReqGetDetail) (res payload.ResGetDetailTeacher, err error)
	GetTeacherOnLesson(ctx context.Context, req lessonPayload.ReqByTeacherId) (res bool, err error)
	SearchTeacher(ctx context.Context, req payload.ReqSearch) (res []payload.ResGetDetailTeacher, err error)
	UpdateTeacher(ctx context.Context, req models.Teacher) (res payload.ResUpdate, err error)
	DeleteTeacher(ctx context.Context, req payload.ReqDelete) (res bool, err error)
}

type ITeacherRepo interface {
	Insert(ctx context.Context, req models.Teacher) (res models.Teacher, err error)
	Get(ctx context.Context, req payload.ReqGetAllTeacher) (res []models.Teacher, err error)
	GetDetail(ctx context.Context, req payload.ReqGetDetail) (res models.Teacher, err error)
	Search(ctx context.Context, req payload.ReqSearch) (res []models.Teacher, err error)
	Update(ctx context.Context, req models.Teacher) (res payload.ResUpdate, err error)
	Delete(ctx context.Context, req payload.ReqDelete) (res bool, err error)
}
