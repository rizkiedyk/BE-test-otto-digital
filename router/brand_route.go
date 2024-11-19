package router

import (
	"test-ottodigital-be/handler"

	"github.com/gin-gonic/gin"
)

func BrandRouter(r *gin.RouterGroup, brandHandler *handler.BrandHandler) {
	brand := r.Group("/brand")
	{
		brand.POST("/", brandHandler.CreateBrand)
	}
}
