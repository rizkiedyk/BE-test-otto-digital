package model

import "github.com/lib/pq"

type Voucher struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	VoucherId   string `json:"voucher_id" gorm:"type:varchar(255);not null;uniqueIndex"`
	BrandID     string `json:"brand_id" gorm:"type:varchar(255);not null"`
	Code        string `json:"code" gorm:"type:varchar(255);not null"`
	CostInPoint int    `json:"cost_in_point" gorm:"not null"`
	CreatedAt   int    `json:"created_at" gorm:"not null"`
	UpdatedAt   int    `json:"updated_at" gorm:"not null"`
	Deleted     bool   `gorm:"default:false" json:"deleted"`
	Brand       Brand  `gorm:"foreignKey:brand_id;references:brand_id"`
}

type VoucherRedemption struct {
	ID            uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	TransactionID string         `json:"transaction_id" gorm:"type:varchar(255);not null;uniqueIndex"`
	TotalPoints   uint           `json:"total_points" gorm:"not null"`
	CreatedAt     int            `json:"created_at" gorm:"not null"`
	UpdatedAt     int            `json:"updated_at" gorm:"not null"`
	Deleted       bool           `gorm:"default:false" json:"deleted"`
	VoucherIDs    pq.StringArray `gorm:"type:text[]" json:"voucher_ids"`
}
