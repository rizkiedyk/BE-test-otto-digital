package service

import (
	"test-ottodigital-be/domain/dto"
	"test-ottodigital-be/domain/model"
	"test-ottodigital-be/repository"
	"test-ottodigital-be/utils"
	"time"

	"github.com/google/uuid"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("main")

type IBrandService interface {
	CreateBrand(brand dto.ReqBrand) error
	GetByID(brandID string) (model.Brand, error)
	GetAll(pagination dto.ReqPagination) (dto.RespBrandGetAll, error)
	UpdateBrand(brand dto.ReqBrand) error
	SoftDelete(brandID string) error
}

type brandService struct {
	repo repository.IBrandRepo
}

func NewBrandService(repo repository.IBrandRepo) IBrandService {
	return &brandService{
		repo: repo,
	}
}

func (s *brandService) CreateBrand(brand dto.ReqBrand) error {
	req := model.Brand{
		BrandID:   uuid.New().String(),
		Name:      brand.Name,
		CreatedAt: int(time.Now().Unix()),
		UpdatedAt: int(time.Now().Unix()),
	}

	err := s.repo.CreateBrand(req)
	if err != nil {
		return err
	}
	return nil
}

func (s *brandService) GetByID(brandID string) (model.Brand, error) {
	brand, err := s.repo.GetByID(brandID)
	if err != nil {
		return model.Brand{}, err
	}

	return brand, nil
}

func (s *brandService) GetAll(reqPage dto.ReqPagination) (dto.RespBrandGetAll, error) {
	pagination := utils.NewPagination(reqPage.Page, reqPage.Limit, 0, reqPage.SortBy, reqPage.SortOrder, reqPage.FilterByKey, reqPage.FilterByValue)

	brands, repoPage, err := s.repo.GetAll(pagination)
	if err != nil {
		return dto.RespBrandGetAll{}, err
	}

	resp := dto.RespBrandGetAll{
		Pagination: repoPage,
		Data:       brands,
	}

	return resp, nil
}

func (s *brandService) UpdateBrand(brand dto.ReqBrand) error {
	existingBrand, err := s.repo.GetByID(brand.BrandID)
	if err != nil {
		return err
	}

	existingBrand.Name = brand.Name
	existingBrand.UpdatedAt = int(time.Now().Unix())

	err = s.repo.UpdateBrand(existingBrand)
	if err != nil {
		return err
	}
	return nil
}

func (s *brandService) SoftDelete(brandID string) error {
	err := s.repo.SoftDelete(brandID)
	if err != nil {
		return err
	}
	return nil
}
