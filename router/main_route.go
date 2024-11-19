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

	brandRepo := repository.NewBrandRepo(db)
	brandService := service.NewBrandService(brandRepo)
	brandHandler := handler.NewBrandHandler(brandService)

	BrandRouter(apiV1, brandHandler)

	return route
}
