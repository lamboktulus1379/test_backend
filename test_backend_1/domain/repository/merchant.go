package repository

import (
	"context"

	"test_backend_1/domain/model"
)

//go:generate mockgen -destination ../../mocks/repositories/mock_imerchant_repository.go -package=mocks test_backend_1/domain/repository IMerchant
type IMerchant interface {
	GetDaily(ctx context.Context, req model.ReqQueryParamMerchant) ([]model.MerchantDetail, int64, error)
	GetDailyOutlet(ctx context.Context, req model.ReqQueryParamMerchant) ([]model.MerchantDetail, int64, error)
}
