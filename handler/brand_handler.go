package handler

import (
	"test-ottodigital-be/domain/dto"
	"test-ottodigital-be/service"

	"github.com/gin-gonic/gin"
)

type BrandHandler struct {
	brandService service.IBrandService
}

func NewBrandHandler(brandService service.IBrandService) *BrandHandler {
	return &BrandHandler{
		brandService: brandService,
	}
}

func (h *BrandHandler) CreateBrand(c *gin.Context) {
	var brand dto.ReqBrand

	if err := c.ShouldBindJSON(&brand); err != nil {
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

	err := h.brandService.CreateBrand(brand)
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

func (h *BrandHandler) GetByID(c *gin.Context) {
	id := c.Param("brand_id")

	brand, err := h.brandService.GetByID(id)
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
		Data: brand,
	})

	return
}
