package model

import "time"

type Merchant struct {
	ID           int64     `gorm:"primaryKey;column:id;type:bigint(20);not null"`
	UserID       int       `gorm:"column:user_id;type:int(40)"`
	MerchantName string    `gorm:"column:merchant_name;type:varchar(40);not null"`
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP"`
	CreatedBy    int64     `gorm:"column:created_by;type:varchar(225);not null"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedBy    int64     `gorm:"column:updated_by;type:varchar(225);not null"`
}

func (Merchant) TableName() string {
	return "Merchants"
}

type ReqQueryParamMerchant struct {
	StartDate  string `json:"start_date" form:"start_date"`
	EndDate    string `json:"end_date" form:"end_date"`
	Month      int    `json:"month" form:"month"`
	PerPage    int    `json:"per_page" form:"per_page"`
	PageNumber int    `json:"page_number" form:"page_number"`
	ID         int    `json:"id"`
}

type MerchantDetail struct {
	MerchantName string    `json:"merchant_name"`
	Omzet        float64   `json:"omzet"`
	CreatedAt    time.Time `json:"created_at"`
	OutletName   string    `json:"outlet_name"`
}
