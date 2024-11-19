package utils

import (
	"test-ottodigital-be/domain/dto"

	"github.com/spf13/viper"
)

func GetEnv(key, defaultValue string) string {
	getEnv := viper.GetString(key)
	if len(getEnv) == 0 || getEnv == "" {
		return defaultValue
	}
	return getEnv

}

func NewPagination(page, limit int, total int64, sortBy, sortOrder, filterByKey, filterByValue string) dto.ReqPagination {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if sortOrder == "" {
		sortOrder = "asc"
	}

	return dto.ReqPagination{
		Page:          page,
		Limit:         limit,
		Total:         total,
		SortBy:        sortBy,
		SortOrder:     sortOrder,
		FilterByKey:   filterByKey,
		FilterByValue: filterByValue,
	}
}

func CalculateOffset(p dto.ReqPagination) int {
	return (p.Page - 1) * p.Limit
}
