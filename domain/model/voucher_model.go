package model

type Voucher struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	BrandID     uint   `json:"brand_id" gorm:"not null"`
	Title       string `json:"title" gorm:"not null"`
	CostInPoint uint   `json:"cost_in_point" gorm:"not null"`
	CreatedAt   int    `json:"created_at" gorm:"not null"`
	UpdatedAt   int    `json:"updated_at" gorm:"not null"`
	Brand       Brand  `json:"brand" gorm:"foreignKey:BrandID"`
}
