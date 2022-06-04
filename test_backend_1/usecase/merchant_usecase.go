package usecase

import (
	"context"
	"log"
	"math"
	"test_backend_1/domain/dto"
	"test_backend_1/domain/model"
	"test_backend_1/domain/repository"
)

type IMerchantUsecase interface {
	ReportDaily(ctx context.Context, req model.ReqQueryParamMerchant) dto.ResMerchant
	ReportDailyOutlet(ctx context.Context, req model.ReqQueryParamMerchant) dto.ResMerchant
}

type MerchantUsecase struct {
	userRepository     repository.IUser
	merchantRepository repository.IMerchant
}

func NewMerchantUsecase(userRepository repository.IUser, merchantRepository repository.IMerchant) IMerchantUsecase {
	return &MerchantUsecase{userRepository: userRepository, merchantRepository: merchantRepository}
}

func (merchantUsecase *MerchantUsecase) ReportDaily(ctx context.Context, req model.ReqQueryParamMerchant) dto.ResMerchant {
	var res dto.ResMerchant

	responses, count, err := merchantUsecase.merchantRepository.GetDaily(ctx, req)
	if err != nil {
		log.Printf("An error occurred: %v", err)
	}

	if count <= 0 {
		return res
	}

	resMerchantDetails := []dto.ResMerchantDetail{}
	for _, response := range responses {
		resMerchantDetail := dto.ResMerchantDetail{
			MerchantName: response.MerchantName,
			Omzet:        response.Omzet,
			OutletName:   response.OutletName,
			CreatedAt:    response.CreatedAt,
		}

		resMerchantDetails = append(resMerchantDetails, resMerchantDetail)
	}
	pagination := dto.Pagination{
		PageNumber:  req.PageNumber,
		PerPage:     req.PerPage,
		TotalPage:   int(math.Ceil(float64(count) / float64(req.PerPage))),
		TotalRecord: int(count),
	}
	res.Data = resMerchantDetails
	res.Pagination = &pagination

	return res
}

func (merchantUsecase *MerchantUsecase) ReportDailyOutlet(ctx context.Context, req model.ReqQueryParamMerchant) dto.ResMerchant {
	var res dto.ResMerchant

	responses, count, err := merchantUsecase.merchantRepository.GetDailyOutlet(ctx, req)
	if err != nil {
		log.Printf("An error occurred: %v", err)
	}

	if count <= 0 {
		return res
	}

	resMerchantDetails := []dto.ResMerchantDetail{}
	for _, response := range responses {
		resMerchantDetail := dto.ResMerchantDetail{
			MerchantName: response.MerchantName,
			Omzet:        response.Omzet,
			OutletName:   response.OutletName,
			CreatedAt:    response.CreatedAt,
		}

		resMerchantDetails = append(resMerchantDetails, resMerchantDetail)
	}
	pagination := dto.Pagination{
		PageNumber:  req.PageNumber,
		PerPage:     req.PerPage,
		TotalPage:   int(math.Ceil(float64(count) / float64(req.PerPage))),
		TotalRecord: int(count),
	}
	res.Data = resMerchantDetails
	res.Pagination = &pagination

	return res
}
