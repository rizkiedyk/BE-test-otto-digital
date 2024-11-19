package repository

import (
	"test-ottodigital-be/domain/model"

	"gorm.io/gorm"
)

type IVoucherRepo interface {
	CreateVoucher(voucher model.Voucher) error
	GetByID(voucherID string) (model.Voucher, error)
	GetByBrandID(brandID string) ([]model.Voucher, error)
}

type voucherRepo struct {
	db *gorm.DB
}

func NewVoucherRepo(db *gorm.DB) IVoucherRepo {
	return &voucherRepo{
		db: db,
	}
}

func (r *voucherRepo) CreateVoucher(voucher model.Voucher) error {
	err := r.db.Create(&voucher).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *voucherRepo) GetByID(voucherID string) (model.Voucher, error) {
	var voucher model.Voucher
	if err := r.db.Preload("Brand").Where("voucher_id = ? AND deleted = false", voucherID).First(&voucher).Error; err != nil {
		return model.Voucher{}, err
	}

	return voucher, nil
}

func (r *voucherRepo) GetByBrandID(brandID string) ([]model.Voucher, error) {
	var vouchers []model.Voucher
	if err := r.db.Preload("Brand").Where("brand_id = ?", brandID).Find(&vouchers).Error; err != nil {
		return nil, err
	}
	return vouchers, nil
}
