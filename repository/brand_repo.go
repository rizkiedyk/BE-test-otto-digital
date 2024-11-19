package repository

import (
	"test-ottodigital-be/domain/model"

	"gorm.io/gorm"
)

type IBrandRepo interface {
	CreateBrand(brand model.Brand) error
	GetByID(brandID string) (model.Brand, error)
}

type brandRepo struct {
	db *gorm.DB
}

func NewBrandRepo(db *gorm.DB) IBrandRepo {
	return &brandRepo{
		db: db,
	}
}

func (r *brandRepo) CreateBrand(brand model.Brand) error {
	err := r.db.Create(&brand).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *brandRepo) GetByID(brandID string) (model.Brand, error) {
	var brand model.Brand
	if err := r.db.Where("brand_id = ? AND deleted = false", brandID).First(&brand).Error; err != nil {
		return model.Brand{}, err
	}
	return brand, nil
}
