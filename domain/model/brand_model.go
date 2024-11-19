package model

import "gorm.io/gorm"

type Brand struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	BrandID   string `json:"brand_id" gorm:"type:varchar(255);not null"`
	Name      string `json:"name" gorm:"not null"`
	CreatedAt int    `json:"created_at" gorm:"not null"`
	UpdatedAt int    `json:"updated_at" gorm:"not null"`
	Deleted   bool   `json:"deleted" gorm:"default:false"`
}

// Hook to set default value for Deleted before creating a brand
func (b *Brand) BeforeCreate(tx *gorm.DB) (err error) {
	if b.Deleted == false {
		b.Deleted = false // Ensure Deleted is set to false
	}
	return nil
}
