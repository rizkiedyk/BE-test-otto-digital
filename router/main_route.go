package router

import (
	"test-ottodigital-be/config"
	"test-ottodigital-be/handler"
	"test-ottodigital-be/repository"
	"test-ottodigital-be/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	route := gin.Default()

	apiV1 := route.Group("/api/v1")

	db := config.ConnectDatabase()

	// Brand
	brandRepo := repository.NewBrandRepo(db)
	brandService := service.NewBrandService(brandRepo)
	brandHandler := handler.NewBrandHandler(brandService)

	// Voucher
	voucherRepo := repository.NewVoucherRepo(db)
	redemptionVoucherRepo := repository.NewRedemptionVoucherRepo(db)
	voucherService := service.NewVoucherService(voucherRepo, redemptionVoucherRepo)
	voucherHandler := handler.NewVoucherHandler(voucherService)

	BrandRouter(apiV1, brandHandler)
	VoucherRouter(apiV1, voucherHandler)

	return route
}
