package students

import (
	"context"
	"gosql/modules/students/models"
	"gosql/modules/students/payload"
)

type IStudentUseCase interface {
	GetAllStudent(ctx context.Context, req payload.ReqGetAllStudents) (res payload.ResGetAllStudent, err error)
	InsertStudent(ctx context.Context, req payload.ReqInsert) (res payload.ReqInsert, err error)
	MassInsertStudent(ctx context.Context, req payload.ReqMassCreate) (res payload.ResMassCreate, err error)
	GetDetail(ctx context.Context, req payload.ReqGetDetail) (res models.Student, err error)
	SearchStudent(ctx context.Context, req payload.ReqSearch) (res []models.Student, err error)
	UpdateStudent(ctx context.Context, req models.Student) (res payload.ResUpdate, err error)
	DeleteStudent(ctx context.Context, req payload.ReqDelete) (res bool, err error)
}

type IStudentRepo interface {
	Insert(ctx context.Context, req models.Student) (res models.Student, err error)
	MassInsert(ctx context.Context, req payload.ReqMassCreate) (res payload.ResMassCreate, err error)
	Get(ctx context.Context, req payload.ReqGetAllStudents) (res payload.ResGetAllStudent, err error)
	GetDetail(ctx context.Context, req payload.ReqGetDetail) (res models.Student, err error)
	Search(ctx context.Context, req payload.ReqSearch) (res []models.Student, err error)
	Update(ctx context.Context, req models.Student) (res payload.ResUpdate, err error)
	Delete(ctx context.Context, req payload.ReqDelete) (res bool, err error)
}
