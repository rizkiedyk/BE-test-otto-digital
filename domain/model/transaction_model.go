package model

type Transaction struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	VoucherID   uint    `json:"voucher_id" gorm:"not null"`
	Quantity    uint    `json:"quantity" gorm:"not null"`
	TotalPoints uint    `json:"total_points" gorm:"not null"`
	Voucher     Voucher `json:"voucher" gorm:"foreignKey:VoucherID"`
	CreatedAt   int     `json:"created_at" gorm:"not null"`
	UpdatedAt   int     `json:"updated_at" gorm:"not null"`
}
