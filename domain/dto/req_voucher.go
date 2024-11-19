package dto

type ReqCreateVoucher struct {
	BrandID     string `json:"brand_id"`
	Code        string `json:"code"`
	CostInPoint int    `json:"cost_in_point"`
}

type ReqRedemption struct {
	VoucherIDs []string `json:"voucher_ids"`
}
