package router

import (
	"test-ottodigital-be/handler"

	"github.com/gin-gonic/gin"
)

func BrandRouter(r *gin.RouterGroup, brandHandler *handler.BrandHandler) {
	brand := r.Group("/brand")
	{
		brand.POST("/", brandHandler.CreateBrand)
		brand.GET("/:brand_id", brandHandler.GetByID)
		brand.GET("/", brandHandler.GetAll)
		brand.PATCH("/:brand_id", brandHandler.UpdateBrand)
		brand.DELETE("/:brand_id", brandHandler.DeleteBrand)
	}
}
