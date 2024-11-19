package router

import (
	"test-ottodigital-be/handler"

	"github.com/gin-gonic/gin"
)

func VoucherRouter(r *gin.RouterGroup, voucherHandler *handler.VoucherHandler) {
	voucher := r.Group("/voucher")
	{
		voucher.POST("/", voucherHandler.CreateVoucher)
		voucher.GET("", voucherHandler.GetVoucherByID)
		voucher.GET("/brand", voucherHandler.GetVoucherByBrandID)
		voucher.POST("/transaction/redemption", voucherHandler.CreateRedemptionVoucher)
		voucher.GET("/transaction/redemption", voucherHandler.GetRedemptionVoucherByID)
	}
}
