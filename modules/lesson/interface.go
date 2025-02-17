package lesson

import (
	"context"
	"gosql/modules/lesson/models"
	"gosql/modules/lesson/payload"
	reportPayload "gosql/modules/report/payload"
)

type ILessonUseCase interface {
	InsertLesson(ctx context.Context, req models.Lesson) (res payload.ResInsert, err error)
	GetLesson(ctx context.Context, req payload.ReqGetAllLesson) (res payload.ResGetAllLesson, err error)
	GetDetail(ctx context.Context, req payload.ReqGetDetail) (res payload.ResGetDetailLesson, err error)
	GetByTeacherID(ctx context.Context, req payload.ReqByTeacherId) (res bool, err error)
	GetByLesson(ctx context.Context, req reportPayload.GetByLesson) (res bool, err error)
	SearchLesson(ctx context.Context, req payload.ReqSearch) (res []payload.ResGetDetailLesson, err error)
	UpdateLesson(ctx context.Context, req models.Lesson) (res payload.ResUpdate, err error)
	DeleteLesson(ctx context.Context, req payload.ReqDelete) (res bool, err error)
}

type ILessonRepo interface {
	Insert(ctx context.Context, req models.Lesson) (res models.Lesson, err error)
	Get(ctx context.Context, req payload.ReqGetAllLesson) (res []models.Lesson, err error)
	GetDetail(ctx context.Context, req payload.ReqGetDetail) (res models.Lesson, err error)
	GetByTeacherID(ctx context.Context, req payload.ReqByTeacherId) (res models.Lesson, err error)
	Search(ctx context.Context, req payload.ReqSearch) (res []models.Lesson, err error)
	Update(ctx context.Context, req models.Lesson) (res payload.ResUpdate, err error)
	Delete(ctx context.Context, req payload.ReqDelete) (res bool, err error)
}
