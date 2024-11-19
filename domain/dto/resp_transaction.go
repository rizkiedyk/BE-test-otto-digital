package dto

import "test-ottodigital-be/domain/model"

type RedemptionResponse struct {
	TransactionID string          `json:"transaction_id"`
	TotalPoints   int             `json:"total_points"`
	Vouchers      []model.Voucher `json:"vouchers"`
}
