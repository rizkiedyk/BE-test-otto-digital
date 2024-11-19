package repository

import (
	"errors"
	"math"
	"test-ottodigital-be/domain/dto"
	"test-ottodigital-be/domain/model"
	"test-ottodigital-be/utils"

	"github.com/op/go-logging"
	"gorm.io/gorm"
)

var logger = logging.MustGetLogger("main")

type IBrandRepo interface {
	CreateBrand(brand model.Brand) error
	GetByID(brandID string) (model.Brand, error)
	GetAll(pagination dto.ReqPagination) ([]model.Brand, dto.Pagination, error)
	UpdateBrand(brand model.Brand) error
	SoftDelete(brandID string) error
}

type brandRepo struct {
	db *gorm.DB
}

func NewBrandRepo(db *gorm.DB) IBrandRepo {
	return &brandRepo{
		db: db,
	}
}

func (r *brandRepo) CreateBrand(brand model.Brand) error {
	err := r.db.Create(&brand).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *brandRepo) GetByID(brandID string) (model.Brand, error) {
	var brand model.Brand
	if err := r.db.Where("brand_id = ? AND deleted = false", brandID).First(&brand).Error; err != nil {
		return model.Brand{}, err
	}
	return brand, nil
}

func (r *brandRepo) GetAll(ReqPagination dto.ReqPagination) ([]model.Brand, dto.Pagination, error) {
	var brands []model.Brand
	var total int64

	query := r.db.Model(&model.Brand{}).Where("deleted = false")

	if ReqPagination.FilterByKey != "" && ReqPagination.FilterByValue != "" {
		query = query.Where(ReqPagination.FilterByKey+" ILIKE ?", "%"+ReqPagination.FilterByValue+"%")
	}

	if ReqPagination.SortBy != "" {
		if ReqPagination.SortOrder == "desc" {
			query = query.Order(ReqPagination.SortBy + " desc")
		} else {
			query = query.Order(ReqPagination.SortBy + " asc")
		}
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, dto.Pagination(ReqPagination), err
	}

	offset := utils.CalculateOffset(ReqPagination)
	query = query.Offset(offset).Limit(ReqPagination.Limit)

	if err := query.Find(&brands).Error; err != nil {
		return nil, dto.Pagination(ReqPagination), err
	}

	ReqPagination.Total = total
	ReqPagination.TotalPage = int(math.Ceil(float64(total) / float64(ReqPagination.Limit)))

	return brands, dto.Pagination(ReqPagination), nil
}

func (r *brandRepo) UpdateBrand(brand model.Brand) error {
	err := r.db.Save(&brand).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *brandRepo) SoftDelete(brandID string) error {
	var brand model.Brand
	if err := r.db.Where("brand_id = ? AND deleted = false", brandID).First(&brand).Error; err != nil {
		return errors.New("brand not found")
	}

	brand.Deleted = true
	if err := r.db.Save(&brand).Error; err != nil {
		return err
	}

	return nil
}
