package test

import (
	"errors"
	"test-ottodigital-be/domain/dto"
	"test-ottodigital-be/domain/model"
	"test-ottodigital-be/service"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockBrandRepository struct {
	mock.Mock
}

func (m *MockBrandRepository) CreateBrand(brand model.Brand) error {
	args := m.Called(brand)
	return args.Error(0)
}

func TestCreateBrand_Success(t *testing.T) {
	mockRepo := new(MockBrandRepository)

	brandService := service.NewBrandService(mockRepo)

	req := dto.ReqBrand{
		Name: "Test Brand",
	}

	expectedBrand := model.Brand{
		BrandID:   uuid.New().String(),
		Name:      req.Name,
		CreatedAt: int(time.Now().Unix()),
		UpdatedAt: int(time.Now().Unix()),
		Deleted:   false,
	}

	mockRepo.On("CreateBrand", mock.MatchedBy(func(brand model.Brand) bool {
		return brand.Name == expectedBrand.Name
	})).Return(nil)

	err := brandService.CreateBrand(req)

	assert.NoError(t, err, "Brand creation should be successful")
	mockRepo.AssertExpectations(t)
}

func TestCreateBrand_Failure(t *testing.T) {
	mockRepo := new(MockBrandRepository)

	brandService := service.NewBrandService(mockRepo)

	req := dto.ReqBrand{
		Name: "",
	}

	expectedBrand := model.Brand{
		BrandID:   uuid.New().String(),
		Name:      req.Name,
		CreatedAt: int(time.Now().Unix()),
		UpdatedAt: int(time.Now().Unix()),
		Deleted:   false,
	}

	expectedError := errors.New("failed to create brand")
	mockRepo.On("CreateBrand", mock.MatchedBy(func(brand model.Brand) bool {
		return brand.Name == expectedBrand.Name
	})).Return(expectedError)

	err := brandService.CreateBrand(req)

	assert.Error(t, err, "Brand creation should fail")
	assert.EqualError(t, err, "failed to create brand", "Error message should match")
	mockRepo.AssertExpectations(t)
}
