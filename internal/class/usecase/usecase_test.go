package usecase_test

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"shiva/shiva-auth/internal/class"
	"shiva/shiva-auth/internal/class/mocks"
	"shiva/shiva-auth/internal/class/usecase"
	"shiva/shiva-auth/utils/baseErrors"
	"testing"
)

var uc class.Usecase
var classRepo mocks.Repository
var classDomain class.Domain
var classListDomain []class.Domain
var uploader s3manager.Uploader

func setup() {
	uc = usecase.NewClassUsecase(&classRepo, &uploader)
	classDomain = class.Domain{
		ID:          1,
		Name:        "ABC",
		IsPasca:     false,
		ImageUrl:    "http://shivalta.com/img/imgtest.jpg",
		ImageHeader: nil,
		Slug:        "test",
	}
	classListDomain = append(classListDomain, classDomain)
}

func TestGetAll(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		setup()
		classRepo.On("GetAll", mock.Anything, mock.Anything).Return(classListDomain, nil).Once()
		data, err := uc.GetAll("", "")

		assert.NoError(t, err)
		assert.Equal(t, data, classListDomain)
	})

	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		classRepo.On("GetAll", mock.Anything, mock.Anything).Return([]class.Domain{}, errors.New("")).Once()
		data, err := uc.GetAll("", "")

		assert.Error(t, err)
		assert.Equal(t, data, []class.Domain{})
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		setup()
		classRepo.On("GetById", mock.Anything, mock.Anything).Return(classDomain, nil).Once()
		data, err := uc.GetById(1)

		assert.NoError(t, err)
		assert.Equal(t, data, classDomain)
	})

	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		classRepo.On("GetById", mock.Anything, mock.Anything).Return(class.Domain{}, errors.New("")).Once()
		data, err := uc.GetById(1)

		assert.Error(t, err)
		assert.Equal(t, data, class.Domain{})
	})

	t.Run("Test Case 3", func(t *testing.T) {
		setup()
		mck := class.Domain{
			ID: 0,
		}
		classRepo.On("GetById", mock.Anything, mock.Anything).Return(mck, nil).Once()
		data, err := uc.GetById(1)

		assert.Error(t, err)
		assert.Equal(t, data, class.Domain{})
	})
}

func TestCreate(t *testing.T) {
	//t.Run("Test Case 2", func(t *testing.T){
	//	setup()
	//	classRepo.On("Create", mock.Anything).Return(classDomain, nil).Once()
	//	data, err := uc.Create(classDomain)
	//
	//	assert.NoError(t, err)
	//	assert.Equal(t, data, classDomain)
	//})
}

func TestUpdate(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		setup()
		classRepo.On("GetById", mock.Anything).Return(class.Domain{}, errors.New("")).Once()
		data, err := uc.Update(classDomain)

		assert.Error(t, err)
		assert.Equal(t, data, class.Domain{})
	})

	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		classRepo.On("GetById", mock.Anything).Return(class.Domain{}, baseErrors.ErrRecordNotFound).Once()
		data, err := uc.Update(classDomain)

		assert.Error(t, err)
		assert.Equal(t, data, class.Domain{})
	})

	t.Run("Test Case 3", func(t *testing.T) {
		setup()
		classRepo.On("GetById", mock.Anything).Return(classDomain, nil).Once()
		classRepo.On("UpdateWithoutImage", mock.Anything).Return(class.Domain{}, errors.New("")).Once()
		data, err := uc.Update(classDomain)

		assert.Error(t, err)
		assert.Equal(t, data, class.Domain{})
	})

	t.Run("Test Case 4", func(t *testing.T) {
		setup()
		classRepo.On("GetById", mock.Anything).Return(classDomain, nil).Once()
		classRepo.On("UpdateWithoutImage", mock.Anything).Return(classDomain, nil).Once()
		data, err := uc.Update(classDomain)

		assert.NoError(t, err)
		assert.Equal(t, data, classDomain)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		setup()
		classRepo.On("GetById", mock.Anything).Return(class.Domain{}, errors.New("")).Once()
		err := uc.Delete(1)

		assert.Error(t, err)
	})

	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		z := class.Domain{
			ID: 0,
		}
		classRepo.On("GetById", mock.Anything).Return(z, nil).Once()
		err := uc.Delete(1)

		assert.Equal(t, err, baseErrors.ErrNotFound)
		assert.Error(t, err)
	})

	t.Run("Test Case 3", func(t *testing.T) {
		setup()
		classRepo.On("GetById", mock.Anything).Return(classDomain, nil).Once()
		classRepo.On("Delete", mock.Anything).Return(errors.New("")).Once()

		err := uc.Delete(1)
		assert.Error(t, err)
	})

	t.Run("Test Case 4", func(t *testing.T) {
		setup()
		classRepo.On("GetById", mock.Anything).Return(classDomain, nil).Once()
		classRepo.On("Delete", mock.Anything).Return(nil).Once()

		err := uc.Delete(1)
		assert.NoError(t, err)
	})
}
