package model

import "gorm.io/gorm"

type Brand struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	BrandID   string `json:"brand_id" gorm:"type:varchar(255);not null;uniqueIndex" migrate:"false"`
	Name      string `json:"name" gorm:"not null"`
	Price     int    `json:"price" gorm:"not null"`
	CreatedAt int    `json:"created_at" gorm:"not null"`
	UpdatedAt int    `json:"updated_at" gorm:"not null"`
	Deleted   bool   `json:"deleted" gorm:"default:false"`
}

func (b *Brand) BeforeCreate(tx *gorm.DB) (err error) {
	if b.Deleted == false {
		b.Deleted = false
	}
	return nil
}
