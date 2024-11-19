package repository

import (
	"test-ottodigital-be/domain/model"

	"gorm.io/gorm"
)

type IRedemptionVoucherRepo interface {
	CreateRedemptionVoucher(redemptionVoucher model.VoucherRedemption) error
	GetTransactionByID(transactionID string) (model.VoucherRedemption, error)
}

type RedemptionVoucherRepo struct {
	db *gorm.DB
}

func NewRedemptionVoucherRepo(db *gorm.DB) IRedemptionVoucherRepo {
	return &RedemptionVoucherRepo{
		db: db,
	}
}

func (r *RedemptionVoucherRepo) CreateRedemptionVoucher(redemptionVoucher model.VoucherRedemption) error {
	err := r.db.Create(&redemptionVoucher).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RedemptionVoucherRepo) GetTransactionByID(transactionID string) (model.VoucherRedemption, error) {
	var transaction model.VoucherRedemption
	if err := r.db.Where("transaction_id = ?", transactionID).First(&transaction).Error; err != nil {
		return model.VoucherRedemption{}, err
	}
	return transaction, nil
}
