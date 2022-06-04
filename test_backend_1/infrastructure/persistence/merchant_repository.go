package persistence

import (
	"context"
	"fmt"
	"log"

	"test_backend_1/domain/model"
	"test_backend_1/domain/repository"

	"gorm.io/gorm"
)

type MerchantRepository struct {
	DB *gorm.DB
}

func NewMerchantRepository(DB *gorm.DB) repository.IMerchant {
	return &MerchantRepository{DB}
}

func (repo *MerchantRepository) GetDaily(ctx context.Context, req model.ReqQueryParamMerchant) ([]model.MerchantDetail, int64, error) {
	var merchantDetails []model.MerchantDetail

	month := req.Month
	userID := req.ID
	var count int64
	queryCount := fmt.Sprintf("SELECT COUNT(*) FROM (SELECT m.merchant_name as merchant_name, SUM(t.bill_total) AS omzet, t.created_at AS created_at FROM Transactions t INNER JOIN Merchants m ON m.id=t.merchant_id INNER JOIN Users u ON u.id=m.user_id WHERE MONTH(t.created_at) = %d AND u.id=%d GROUP BY DATE(t.created_at), t.merchant_id ORDER BY t.created_at DESC) result", month, userID)

	query := fmt.Sprintf("SELECT m.merchant_name as merchant_name, SUM(t.bill_total) AS omzet, t.created_at AS created_at FROM Transactions t INNER JOIN Merchants m ON m.id=t.merchant_id INNER JOIN Users u ON u.id=m.user_id WHERE MONTH(t.created_at) = %d AND u.id=%d GROUP BY DATE(t.created_at), t.merchant_id ORDER BY t.created_at DESC LIMIT %d OFFSET %d", month, userID, req.PerPage, (req.PageNumber-1)*req.PerPage)

	if err := repo.DB.Debug().WithContext(ctx).Raw(queryCount).Scan(&count).Error; err != nil {
		log.Printf("Failed Get User With Error : %v", err)
		return merchantDetails, 0, err
	}
	if err := repo.DB.Debug().WithContext(ctx).Raw(query).Find(&merchantDetails).Error; err != nil {
		log.Printf("Failed Get User With Error : %v", err)
		return merchantDetails, 0, err
	}

	return merchantDetails, count, nil
}

func (repo *MerchantRepository) GetDailyOutlet(ctx context.Context, req model.ReqQueryParamMerchant) ([]model.MerchantDetail, int64, error) {
	var merchantDetails []model.MerchantDetail

	month := req.Month
	userID := req.ID
	var count int64
	queryCount := fmt.Sprintf("SELECT COUNT(*) FROM (SELECT m.merchant_name as merchant_name, o.outlet_name as outlet_name, SUM(t.bill_total) AS omzet, t.created_at AS created_at FROM Transactions t INNER JOIN Merchants m ON m.id=t.merchant_id INNER JOIN Users u ON u.id=m.user_id INNER JOIN Outlets o ON o.merchant_id=m.id WHERE MONTH(t.created_at) = %d AND u.id=%d GROUP BY DATE(t.created_at), t.outlet_id ORDER BY t.created_at DESC) result", month, userID)

	query := fmt.Sprintf("SELECT m.merchant_name as merchant_name, o.outlet_name as outlet_name, SUM(t.bill_total) AS omzet, t.created_at AS created_at FROM Transactions t INNER JOIN Merchants m ON m.id=t.merchant_id INNER JOIN Users u ON u.id=m.user_id INNER JOIN Outlets o ON o.merchant_id=m.id WHERE MONTH(t.created_at) = %d AND u.id=%d GROUP BY DATE(t.created_at), t.outlet_id ORDER BY t.created_at DESC LIMIT %d OFFSET %d", month, userID, req.PerPage, (req.PageNumber-1)*req.PerPage)

	if err := repo.DB.Debug().WithContext(ctx).Raw(queryCount).Scan(&count).Error; err != nil {
		log.Printf("Failed Get User With Error : %v", err)
		return merchantDetails, 0, err
	}
	if err := repo.DB.Debug().WithContext(ctx).Raw(query).Find(&merchantDetails).Error; err != nil {
		log.Printf("Failed Get User With Error : %v", err)
		return merchantDetails, 0, err
	}

	return merchantDetails, count, nil
}
