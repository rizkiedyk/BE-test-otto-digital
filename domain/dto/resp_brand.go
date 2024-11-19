package dto

import (
	"test-ottodigital-be/domain/model"
)

type RespBrandGetAll struct {
	Pagination Pagination    `json:"pagination"`
	Data       []model.Brand `json:"data"`
}
