package service

import (
	"test-ottodigital-be/domain/dto"
	"test-ottodigital-be/domain/model"
	"test-ottodigital-be/repository"
	"time"

	"github.com/google/uuid"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("main")

type IBrandService interface {
	CreateBrand(brand dto.ReqBrand) error
	GetByID(brandID string) (model.Brand, error)
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
