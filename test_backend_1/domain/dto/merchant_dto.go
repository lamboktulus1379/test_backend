package dto

import "time"

type ResMerchant struct {
	*Pagination
	Data []ResMerchantDetail `json:"data"`
}

type ResMerchantDetail struct {
	MerchantName string    `json:"merchant_name,omitempty"`
	Omzet        float64   `json:"omzet,omitempty"`
	OutletName   string    `json:"outlet_name,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}
