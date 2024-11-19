package handler

import (
	"test-ottodigital-be/domain/dto"
	"test-ottodigital-be/service"

	"github.com/gin-gonic/gin"
)

type VoucherHandler struct {
	voucherService service.IVoucherService
}

func NewVoucherHandler(voucherService service.IVoucherService) *VoucherHandler {
	return &VoucherHandler{
		voucherService: voucherService,
	}
}

func (v *VoucherHandler) CreateVoucher(c *gin.Context) {
	var reqVoucher dto.ReqCreateVoucher
	if err := c.ShouldBindJSON(&reqVoucher); err != nil {
		c.JSON(400, dto.Resp{
			Meta: dto.Meta{
				Success: false,
				Code:    400,
				Message: "Bad Request",
				ErrorDetail: map[string]interface{}{
					"error": err.Error(),
				},
			},
			Data: nil,
		})
		return
	}

	err := v.voucherService.CreateVoucher(reqVoucher)
	if err != nil {
		c.JSON(400, dto.Resp{
			Meta: dto.Meta{
				Success:     false,
				Code:        400,
				Message:     "Bad Request",
				ErrorDetail: err.Error(),
			},
			Data: nil,
		})
		return
	}
	c.JSON(200, dto.Resp{
		Meta: dto.Meta{
			Success:     true,
			Code:        200,
			Message:     "Success",
			ErrorDetail: nil,
		},
		Data: nil,
	})
}

func (v *VoucherHandler) GetVoucherByID(c *gin.Context) {
	voucherID := c.Query("id")

	voucher, err := v.voucherService.GetByID(voucherID)
	if err != nil {
		c.JSON(400, dto.Resp{
			Meta: dto.Meta{
				Success:     false,
				Code:        400,
				Message:     "Bad Request",
				ErrorDetail: err.Error(),
			},
			Data: nil,
		})
		return
	}

	c.JSON(200, dto.Resp{
		Meta: dto.Meta{
			Success:     true,
			Code:        200,
			Message:     "Success",
			ErrorDetail: nil,
		},
		Data: voucher,
	})
	return
}

func (v *VoucherHandler) GetVoucherByBrandID(c *gin.Context) {
	brandID := c.Query("id")

	vouchers, err := v.voucherService.GetByBrandID(brandID)
	if err != nil {
		c.JSON(400, dto.Resp{
			Meta: dto.Meta{
				Success:     false,
				Code:        400,
				Message:     "Bad Request",
				ErrorDetail: err.Error(),
			},
			Data: nil,
		})
		return
	}

	c.JSON(200, dto.Resp{
		Meta: dto.Meta{
			Success:     true,
			Code:        200,
			Message:     "Success",
			ErrorDetail: nil,
		},
		Data: vouchers,
	})
	return
}

func (v *VoucherHandler) CreateRedemptionVoucher(c *gin.Context) {
	var reqRedemption dto.ReqRedemption
	if err := c.ShouldBindJSON(&reqRedemption); err != nil {
		c.JSON(400, dto.Resp{
			Meta: dto.Meta{
				Success:     false,
				Code:        400,
				Message:     "Bad Request",
				ErrorDetail: err.Error(),
			},
			Data: nil,
		})
		return
	}

	err := v.voucherService.CreateRedemptionVoucher(reqRedemption)
	if err != nil {
		c.JSON(400, dto.Resp{
			Meta: dto.Meta{
				Success:     false,
				Code:        400,
				Message:     "Bad Request",
				ErrorDetail: err.Error(),
			},
			Data: nil,
		})
		return
	}

	c.JSON(200, dto.Resp{
		Meta: dto.Meta{
			Success:     true,
			Code:        200,
			Message:     "Success",
			ErrorDetail: nil,
		},
		Data: nil,
	})
	return
}

func (v *VoucherHandler) GetRedemptionVoucherByID(c *gin.Context) {
	transactionID := c.Query("id")

	transaction, err := v.voucherService.GetRedemptionVoucher(transactionID)
	if err != nil {
		c.JSON(400, dto.Resp{
			Meta: dto.Meta{
				Success:     false,
				Code:        400,
				Message:     "Bad Request",
				ErrorDetail: err.Error(),
			},
			Data: nil,
		})
		return
	}

	c.JSON(200, dto.Resp{
		Meta: dto.Meta{
			Success:     true,
			Code:        200,
			Message:     "Success",
			ErrorDetail: nil,
		},
		Data: transaction,
	})
	return
}
