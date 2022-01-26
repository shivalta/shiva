package usecase_test

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"shiva/shiva-auth/internal/categories"
	"shiva/shiva-auth/internal/categories/mocks"
	"shiva/shiva-auth/internal/categories/usecase"
	"shiva/shiva-auth/internal/class"
	_mck_class "shiva/shiva-auth/internal/class/mocks"
	"shiva/shiva-auth/utils/baseErrors"
	"testing"
)

var uc categories.Usecase
var ucClass _mck_class.Usecase
var categoriesRepo mocks.Repository
var categoriesDomain categories.Domain
var categoriesListDomain []categories.Domain
var uploader s3manager.Uploader

func setup() {
	uc = usecase.NewCategoriesUsecase(&categoriesRepo, &uploader, &ucClass)
	categoriesDomain = categories.Domain{
		ID:             1,
		ProductClassId: 0,
		ProductClass: categories.ProductClass{
			ID:          0,
			Name:        "",
			IsPasca:     false,
			ImageUrl:    "",
			ImageHeader: nil,
			Slug:        "",
		},
		Name:        "ABC",
		ImageUrl:    "http://shivalta.com/img/imgtest.jpg",
		ImageHeader: nil,
		Slug:        "test",
		Tax:         0,
	}
	categoriesListDomain = append(categoriesListDomain, categoriesDomain)
}

func TestGetAll(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		setup()
		categoriesRepo.On("GetAll", mock.Anything, mock.Anything).Return(categoriesListDomain, nil).Once()
		data, err := uc.GetAll("", "")

		assert.NoError(t, err)
		assert.Equal(t, data, categoriesListDomain)
	})

	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		categoriesRepo.On("GetAll", mock.Anything, mock.Anything).Return([]categories.Domain{}, errors.New("")).Once()
		data, err := uc.GetAll("", "")

		assert.Error(t, err)
		assert.Equal(t, data, []categories.Domain{})
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		setup()
		categoriesRepo.On("GetById", mock.Anything, mock.Anything).Return(categoriesDomain, nil).Once()
		data, err := uc.GetById(1)

		assert.NoError(t, err)
		assert.Equal(t, data, categoriesDomain)
	})

	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		categoriesRepo.On("GetById", mock.Anything, mock.Anything).Return(categories.Domain{}, errors.New("")).Once()
		data, err := uc.GetById(1)

		assert.Error(t, err)
		assert.Equal(t, data, categories.Domain{})
	})

	t.Run("Test Case 3", func(t *testing.T) {
		setup()
		mck := categories.Domain{
			ID: 0,
		}
		categoriesRepo.On("GetById", mock.Anything, mock.Anything).Return(mck, nil).Once()
		data, err := uc.GetById(1)

		assert.Error(t, err)
		assert.Equal(t, data, categories.Domain{})
	})
}

func TestCreate(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		setup()
		ucClass.On("GetById", mock.Anything).Return(class.Domain{}, errors.New("")).Once()

		data, err := uc.Create(categoriesDomain)

		assert.Error(t, err)
		assert.Equal(t, data, categories.Domain{})
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		setup()
		categoriesRepo.On("GetById", mock.Anything).Return(categories.Domain{}, errors.New("")).Once()
		data, err := uc.Update(categoriesDomain)

		assert.Error(t, err)
		assert.Equal(t, data, categories.Domain{})
	})

	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		categoriesRepo.On("GetById", mock.Anything).Return(categories.Domain{}, baseErrors.ErrRecordNotFound).Once()
		data, err := uc.Update(categoriesDomain)

		assert.Error(t, err)
		assert.Equal(t, data, categories.Domain{})
	})

	t.Run("Test Case 3", func(t *testing.T) {
		setup()
		categoriesRepo.On("GetById", mock.Anything).Return(categoriesDomain, nil).Once()
		categoriesRepo.On("UpdateWithoutImage", mock.Anything).Return(categories.Domain{}, errors.New("")).Once()
		data, err := uc.Update(categoriesDomain)

		assert.Error(t, err)
		assert.Equal(t, data, categories.Domain{})
	})

	t.Run("Test Case 4", func(t *testing.T) {
		setup()
		categoriesRepo.On("GetById", mock.Anything).Return(categoriesDomain, nil).Once()
		categoriesRepo.On("UpdateWithoutImage", mock.Anything).Return(categories.Domain{}, errors.New("")).Once()
		data, err := uc.Update(categoriesDomain)

		assert.Error(t, err)
		assert.Equal(t, data, categories.Domain{})
	})

	t.Run("Test Case 5", func(t *testing.T) {
		setup()
		categoriesRepo.On("GetById", mock.Anything).Return(categoriesDomain, nil).Once()
		categoriesRepo.On("UpdateWithoutImage", mock.Anything).Return(categoriesDomain, nil).Once()
		data, err := uc.Update(categoriesDomain)

		assert.Nil(t, err)
		assert.Equal(t, data, categoriesDomain)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		setup()
		categoriesRepo.On("GetById", mock.Anything).Return(categories.Domain{}, errors.New("")).Once()
		err := uc.Delete(1)

		assert.Error(t, err)
	})

	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		z := categories.Domain{
			ID: 0,
		}
		categoriesRepo.On("GetById", mock.Anything).Return(z, nil).Once()
		err := uc.Delete(1)

		assert.Equal(t, err, baseErrors.ErrNotFound)
		assert.Error(t, err)
	})

	t.Run("Test Case 3", func(t *testing.T) {
		setup()
		categoriesRepo.On("GetById", mock.Anything).Return(categoriesDomain, nil).Once()
		categoriesRepo.On("Delete", mock.Anything).Return(errors.New("")).Once()

		err := uc.Delete(1)
		assert.Error(t, err)
	})

	t.Run("Test Case 4", func(t *testing.T) {
		setup()
		categoriesRepo.On("GetById", mock.Anything).Return(categoriesDomain, nil).Once()
		categoriesRepo.On("Delete", mock.Anything).Return(nil).Once()

		err := uc.Delete(1)
		assert.NoError(t, err)
	})
}
