package service

import (
	"test-ottodigital-be/domain/dto"
	"test-ottodigital-be/domain/model"
	"test-ottodigital-be/repository"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type IVoucherService interface {
	CreateVoucher(reqVoucher dto.ReqCreateVoucher) error
	GetByID(voucherID string) (model.Voucher, error)
	GetByBrandID(brandID string) ([]model.Voucher, error)
	CreateRedemptionVoucher(redemptionVoucher dto.ReqRedemption) error
	GetRedemptionVoucher(transactionID string) (dto.RedemptionResponse, error)
}

type voucherService struct {
	repo           repository.IVoucherRepo
	repoRedemption repository.IRedemptionVoucherRepo
}

func NewVoucherService(repo repository.IVoucherRepo, repoRedemption repository.IRedemptionVoucherRepo) IVoucherService {
	return &voucherService{
		repo:           repo,
		repoRedemption: repoRedemption,
	}
}

func (s *voucherService) CreateVoucher(reqVoucher dto.ReqCreateVoucher) error {
	req := model.Voucher{
		VoucherId:   uuid.New().String(),
		BrandID:     reqVoucher.BrandID,
		Code:        reqVoucher.Code,
		CostInPoint: reqVoucher.CostInPoint,
		CreatedAt:   int(time.Now().Unix()),
		UpdatedAt:   int(time.Now().Unix()),
	}

	err := s.repo.CreateVoucher(req)
	if err != nil {
		return err
	}
	return nil
}

func (s *voucherService) GetByID(voucherID string) (model.Voucher, error) {
	voucher, err := s.repo.GetByID(voucherID)
	if err != nil {
		return model.Voucher{}, err
	}
	return voucher, nil
}

func (s *voucherService) GetByBrandID(brandID string) ([]model.Voucher, error) {
	vouchers, err := s.repo.GetByBrandID(brandID)
	if err != nil {
		return nil, err
	}
	return vouchers, nil
}

func (s *voucherService) CreateRedemptionVoucher(redemptionVoucher dto.ReqRedemption) error {
	var totalPoints int
	for _, voucherID := range redemptionVoucher.VoucherIDs {
		voucher, err := s.repo.GetByID(voucherID)
		if err != nil {
			return err
		}

		totalPoints += voucher.CostInPoint
	}

	req := model.VoucherRedemption{
		TransactionID: uuid.New().String(),
		TotalPoints:   uint(totalPoints),
		CreatedAt:     int(time.Now().Unix()),
		UpdatedAt:     int(time.Now().Unix()),
		VoucherIDs:    pq.StringArray(redemptionVoucher.VoucherIDs),
	}

	err := s.repoRedemption.CreateRedemptionVoucher(req)
	if err != nil {
		return err
	}
	return nil
}

func (s *voucherService) GetRedemptionVoucher(transactionID string) (dto.RedemptionResponse, error) {
	transaction, err := s.repoRedemption.GetTransactionByID(transactionID)
	if err != nil {
		return dto.RedemptionResponse{}, err
	}

	var vouchers []model.Voucher
	for _, voucherID := range transaction.VoucherIDs {
		voucher, err := s.repo.GetByID(voucherID)
		if err != nil {
			return dto.RedemptionResponse{}, err
		}
		vouchers = append(vouchers, voucher)
	}

	resp := dto.RedemptionResponse{
		TransactionID: transaction.TransactionID,
		TotalPoints:   int(transaction.TotalPoints),
		Vouchers:      vouchers,
	}

	return resp, nil
}
