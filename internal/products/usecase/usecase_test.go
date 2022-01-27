package usecase_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"shiva/shiva-auth/internal/categories"
	catmock "shiva/shiva-auth/internal/categories/mocks"
	"shiva/shiva-auth/internal/class"
	clsmock "shiva/shiva-auth/internal/class/mocks"
	"shiva/shiva-auth/internal/products"
	"shiva/shiva-auth/internal/products/mocks"
	"shiva/shiva-auth/internal/products/usecase"
	"shiva/shiva-auth/utils/baseErrors"
	"testing"
)

var classUsecase clsmock.Usecase
var categoryUsecase catmock.Usecase
var uc products.Usecase
var productRepo mocks.Repository
var productDomain products.Domain
var productListDomain []products.Domain
var categoryDomain categories.Domain
var classDomain class.Domain

func setup() {
	uc = usecase.NewProductsUsecase(&productRepo, &classUsecase, &categoryUsecase)
	productDomain = products.Domain{
		ID:                1,
		ProductClassId:    1,
		ProductClass:      products.Class{},
		ProductCategoryId: 1,
		ProductCategory:   products.Categories{},
		Sku:               "ABCC",
		Name:              "ABC",
		AdminFee:          100,
		Stock:             1000,
	}
	categoryDomain = categories.Domain{
		ID:             10,
		ProductClassId: 1,
		Name:           "ABC",
		ImageUrl:       "ACA",
		Slug:           "z",
		Tax:            10,
	}
	classDomain = class.Domain{
		ID:       1,
		Name:     "ABC",
		IsPasca:  false,
		ImageUrl: "BAA",
		Slug:     "zz",
	}
	productListDomain = append(productListDomain, productDomain)
}

func TestGetAll(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		setup()
		productRepo.On("GetAll", mock.Anything, mock.Anything).Return(productListDomain, nil).Once()
		data, err := uc.GetAll("", "")

		assert.NoError(t, err)
		assert.Equal(t, data, productListDomain)
	})

	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		productRepo.On("GetAll", mock.Anything, mock.Anything).Return([]products.Domain{}, errors.New("")).Once()
		data, err := uc.GetAll("", "")

		assert.Error(t, err)
		assert.Equal(t, data, []products.Domain{})
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		setup()
		productRepo.On("GetById", mock.Anything, mock.Anything).Return(productDomain, nil).Once()
		data, err := uc.GetById(1)

		assert.NoError(t, err)
		assert.Equal(t, data, productDomain)
	})

	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		productRepo.On("GetById", mock.Anything, mock.Anything).Return(products.Domain{}, errors.New("")).Once()
		data, err := uc.GetById(1)

		assert.Error(t, err)
		assert.Equal(t, data, products.Domain{})
	})

	t.Run("Test Case 3", func(t *testing.T) {
		setup()
		mck := products.Domain{
			ID: 0,
		}
		productRepo.On("GetById", mock.Anything, mock.Anything).Return(mck, nil).Once()
		data, err := uc.GetById(1)

		assert.Error(t, err)
		assert.Equal(t, data, products.Domain{})
	})
}

func TestCreate(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		setup()
		categoryUsecase.On("GetById", mock.Anything).Return(categories.Domain{}, errors.New("")).Once()
		data, err := uc.Create(productDomain)

		assert.Error(t, err)
		assert.Equal(t, data, products.Domain{})
	})
	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		categoryDomain.ProductClassId = 5
		categoryUsecase.On("GetById", mock.Anything).Return(categoryDomain, nil).Once()
		data, err := uc.Create(productDomain)

		assert.Equal(t, err, baseErrors.ErrProductClassIdNotSync)
		assert.Equal(t, data, products.Domain{})
	})
	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		categoryUsecase.On("GetById", mock.Anything).Return(categoryDomain, nil).Once()
		classUsecase.On("GetById", mock.Anything).Return(class.Domain{}, errors.New("")).Once()
		data, err := uc.Create(productDomain)

		assert.Error(t, err)
		assert.Equal(t, data, products.Domain{})
	})
	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		categoryUsecase.On("GetById", mock.Anything).Return(categoryDomain, nil).Once()
		classUsecase.On("GetById", mock.Anything).Return(classDomain, nil).Once()
		productRepo.On("Create", mock.Anything).Return(products.Domain{}, errors.New("")).Once()
		data, err := uc.Create(productDomain)

		assert.Error(t, err)
		assert.Equal(t, data, products.Domain{})
	})
	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		productDomain.ProductClass = products.Class{
			ID:       classDomain.ID,
			Name:     classDomain.Name,
			IsPasca:  classDomain.IsPasca,
			ImageUrl: classDomain.ImageUrl,
			Slug:     classDomain.Slug,
		}
		productDomain.ProductCategory = products.Categories{
			ID:             categoryDomain.ID,
			ProductClassId: categoryDomain.ProductClassId,
			Name:           categoryDomain.Name,
			ImageUrl:       categoryDomain.ImageUrl,
			Slug:           categoryDomain.Slug,
			Tax:            categoryDomain.Tax,
		}
		categoryUsecase.On("GetById", mock.Anything).Return(categoryDomain, nil).Once()
		classUsecase.On("GetById", mock.Anything).Return(classDomain, nil).Once()
		productRepo.On("Create", mock.Anything).Return(productDomain, nil).Once()
		data, err := uc.Create(productDomain)

		assert.NoError(t, err)
		assert.Equal(t, data, productDomain)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		setup()
		productRepo.On("GetById", mock.Anything).Return(products.Domain{}, errors.New("")).Once()
		data, err := uc.Update(productDomain)

		assert.Error(t, err)
		assert.Equal(t, data, products.Domain{})
	})

	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		productRepo.On("GetById", mock.Anything).Return(productDomain, nil).Once()
		productRepo.On("UpdateWithoutImage", mock.Anything).Return(products.Domain{}, errors.New("")).Once()
		data, err := uc.Update(productDomain)

		assert.Error(t, err)
		assert.Equal(t, data, products.Domain{})
	})

	t.Run("Test Case 3", func(t *testing.T) {
		setup()
		productRepo.On("GetById", mock.Anything).Return(productDomain, nil).Once()
		productRepo.On("UpdateWithoutImage", mock.Anything).Return(productDomain, nil).Once()
		data, err := uc.Update(productDomain)

		assert.NoError(t, err)
		assert.Equal(t, data, productDomain)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		setup()
		productRepo.On("GetById", mock.Anything).Return(products.Domain{}, errors.New("")).Once()
		err := uc.Delete(1)

		assert.Error(t, err)
	})

	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		z := products.Domain{
			ID: 0,
		}
		productRepo.On("GetById", mock.Anything).Return(z, nil).Once()
		err := uc.Delete(1)

		assert.Equal(t, err, baseErrors.ErrNotFound)
		assert.Error(t, err)
	})

	t.Run("Test Case 3", func(t *testing.T) {
		setup()
		productRepo.On("GetById", mock.Anything).Return(productDomain, nil).Once()
		productRepo.On("Delete", mock.Anything).Return(errors.New("")).Once()

		err := uc.Delete(1)
		assert.Error(t, err)
	})

	t.Run("Test Case 4", func(t *testing.T) {
		setup()
		productRepo.On("GetById", mock.Anything).Return(productDomain, nil).Once()
		productRepo.On("Delete", mock.Anything).Return(nil).Once()

		err := uc.Delete(1)
		assert.NoError(t, err)
	})
}
