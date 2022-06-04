package model

import "time"

type Transaction struct {
	ID         int64     `gorm:"primaryKey;column:id;type:bigint(20);not null"`
	MerchantID int64     `gorm:"column:merchant_id;type:bigint(20);not null"`
	OutletID   int64     `gorm:"column:outlet_id;type:bigint(20);not null"`
	BillTotal  float64   `gorm:"column:bill_total;not null"`
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP"`
	CreatedBy  int64     `gorm:"column:created_by;type:varchar(225);not null"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedBy  int64     `gorm:"column:updated_by;type:varchar(225);not null"`
}

func (Transaction) TableName() string {
	return "Transactions"
}
