package usecase

import (
	"context"
	"gosql/modules/report/payload"
)

func (s *reportUseCase) DeleteReport(ctx context.Context, req payload.ReqDelete) (res bool, err error) {

	// Update
	if res, err = s.reportEntity.ReportRepo.Delete(ctx, req); err != nil {
		return res, err
	}

	return res, err
}
