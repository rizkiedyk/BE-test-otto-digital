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

func (m *MockBrandRepository) GetByID(brandID string) (model.Brand, error) {
	args := m.Called(brandID)
	return args.Get(0).(model.Brand), args.Error(1)
}

func (m *MockBrandRepository) GetAll(pagination dto.ReqPagination) ([]model.Brand, dto.Pagination, error) {
	args := m.Called(pagination)
	return args.Get(0).([]model.Brand), args.Get(1).(dto.Pagination), args.Error(2)
}

func (m *MockBrandRepository) UpdateBrand(brand model.Brand) error {
	args := m.Called(brand)
	return args.Error(0)
}

func (m *MockBrandRepository) SoftDelete(brandID string) error {
	args := m.Called(brandID)
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

func TestGetByID_Success(t *testing.T) {
	mockRepo := new(MockBrandRepository)

	brandService := service.NewBrandService(mockRepo)

	brandID := uuid.New().String()

	expectedBrand := model.Brand{
		BrandID: brandID,
		Name:    "Test Brand",
	}

	mockRepo.On("GetByID", brandID).Return(expectedBrand, nil)

	brand, err := brandService.GetByID(brandID)

	assert.NoError(t, err)
	assert.Equal(t, expectedBrand, brand)
	mockRepo.AssertExpectations(t)
}

func TestGetByID_Failure(t *testing.T) {
	mockRepo := new(MockBrandRepository)

	brandService := service.NewBrandService(mockRepo)

	brandID := uuid.New().String()

	mockRepo.On("GetByID", brandID).Return(model.Brand{}, errors.New("brand not found"))

	brand, err := brandService.GetByID(brandID)

	assert.Error(t, err)
	assert.EqualError(t, err, "brand not found")
	assert.Equal(t, model.Brand{}, brand)
	mockRepo.AssertExpectations(t)
}

func TestGetAll_Success(t *testing.T) {
	mockRepo := new(MockBrandRepository)
	mockPagination := dto.ReqPagination{
		Page:          1,
		Limit:         10,
		Total:         2,
		SortBy:        "name",
		SortOrder:     "asc",
		FilterByKey:   "name",
		FilterByValue: "Brand",
	}

	expectedBrands := []model.Brand{
		{BrandID: uuid.New().String(), Name: "Brand 1"},
		{BrandID: uuid.New().String(), Name: "Brand 2"},
	}

	expectedPagination := dto.Pagination{
		Page:      1,
		Limit:     10,
		Total:     2,
		SortBy:    "name",
		SortOrder: "asc",
		TotalPage: 1,
	}

	mockRepo.On("GetAll", mock.MatchedBy(func(arg dto.ReqPagination) bool {
		return arg.Total == 0
	})).Return(expectedBrands, expectedPagination, nil)

	brandService := service.NewBrandService(mockRepo)

	resp, err := brandService.GetAll(mockPagination)

	assert.NoError(t, err, "There should be no error")
	assert.Equal(t, expectedPagination.Page, resp.Pagination.Page, "Pagination Page should match")
	assert.Equal(t, expectedPagination.Limit, resp.Pagination.Limit, "Pagination Limit should match")
	assert.Equal(t, expectedPagination.Total, resp.Pagination.Total, "Pagination Total should match")
	assert.Equal(t, expectedPagination.SortBy, resp.Pagination.SortBy, "Pagination SortBy should match")
	assert.Equal(t, expectedPagination.SortOrder, resp.Pagination.SortOrder, "Pagination SortOrder should match")
	assert.Equal(t, len(expectedBrands), len(resp.Data), "Brand data length should match")
	assert.Equal(t, expectedBrands[0].Name, resp.Data[0].Name, "First brand name should match")
	mockRepo.AssertExpectations(t)
}

func TestGetAll_Failure(t *testing.T) {

	mockRepo := new(MockBrandRepository)
	mockPagination := dto.ReqPagination{
		Page:          1,
		Limit:         10,
		Total:         0,
		SortBy:        "name",
		SortOrder:     "asc",
		FilterByKey:   "brand_name",
		FilterByValue: "test",
	}
	expectedBrands := []model.Brand{
		{ID: 1, Name: "Brand 1"},
		{ID: 2, Name: "Brand 2"},
	}
	expectedPagination := dto.Pagination{
		Page:      1,
		Limit:     10,
		Total:     2,
		SortBy:    "name",
		SortOrder: "asc",
	}

	mockRepo.On("GetAll", mockPagination).Return(expectedBrands, expectedPagination, nil)

	brandService := service.NewBrandService(mockRepo)

	resp, err := brandService.GetAll(mockPagination)

	assert.NoError(t, err)
	assert.Equal(t, expectedPagination.Page, resp.Pagination.Page, "Pagination Page should match")
	assert.Equal(t, expectedPagination.Limit, resp.Pagination.Limit, "Pagination Limit should match")
	assert.Equal(t, expectedPagination.Total, resp.Pagination.Total, "Pagination Total should match")
	assert.Equal(t, expectedPagination.SortBy, resp.Pagination.SortBy, "Pagination SortBy should match")
	assert.Equal(t, expectedPagination.SortOrder, resp.Pagination.SortOrder, "Pagination SortOrder should match")
	assert.Equal(t, len(expectedBrands), len(resp.Data), "Brand data length should match")
	assert.Equal(t, expectedBrands[0].Name, resp.Data[0].Name, "First brand name should match")
	mockRepo.AssertExpectations(t)
}

func TestGetAll_Error(t *testing.T) {

	mockRepo := new(MockBrandRepository)
	mockPagination := dto.ReqPagination{
		Page:          1,
		Limit:         10,
		Total:         0,
		SortBy:        "name",
		SortOrder:     "asc",
		FilterByKey:   "brand_name",
		FilterByValue: "test",
	}

	mockRepo.On("GetAll", mockPagination).Return([]model.Brand{}, dto.Pagination{}, errors.New("database error"))

	brandService := service.NewBrandService(mockRepo)

	resp, err := brandService.GetAll(mockPagination)

	assert.Error(t, err)
	assert.Empty(t, resp.Data, "Data should be empty on error")
	mockRepo.AssertExpectations(t)
}

func TestUpdateBrand_Success(t *testing.T) {
	mockRepo := new(MockBrandRepository)
	brandService := service.NewBrandService(mockRepo)

	brand := dto.ReqBrand{
		BrandID: uuid.New().String(),
		Name:    "Updated Brand Name",
	}

	mockRepo.On("GetByID", brand.BrandID).Return(model.Brand{
		BrandID: brand.BrandID,
		Name:    "Old Brand Name",
	}, nil)

	mockRepo.On("UpdateBrand", mock.MatchedBy(func(updatedBrand model.Brand) bool {
		return updatedBrand.Name == brand.Name
	})).Return(nil)

	err := brandService.UpdateBrand(brand)

	assert.NoError(t, err, "Brand update should be successful")

	mockRepo.AssertExpectations(t)
}
