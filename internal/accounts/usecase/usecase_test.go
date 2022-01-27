package usecase_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"shiva/shiva-auth/cmd/http/middlewares"
	"shiva/shiva-auth/internal/accounts"
	"shiva/shiva-auth/internal/accounts/mocks"
	"shiva/shiva-auth/internal/accounts/usecase"
	"testing"
)

var uc accounts.Usecase
var userRepo mocks.Repository
var userDomain accounts.Domain
var userListDomain []accounts.Domain
var bcryptPassword string
var token string

func setup() {
	uc = usecase.NewAccountUsecase(&userRepo, &middlewares.ConfigJWT{})
	bcryptPassword = "$2a$12$JsWQHjj8nEGs5BPCW5ExtuMApPZSAESqbYZeeOBuNs7wNo5Pfv3vC"
	token = "xuxu"
	userDomain = accounts.Domain{
		ID:         1,
		Name:       "ABC",
		Email:      "shiva@shiva.co",
		Handphone:  "08123",
		Address:    "JL ABC",
		IsAdmin:    false,
		IsActive:   false,
		Password:   "123",
		Repassword: "123",
	}
	userListDomain = append(userListDomain, userDomain)
}

func TestGetAll(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		setup()
		userRepo.On("GetAll", mock.Anything, mock.Anything).Return(userListDomain, nil).Once()
		data, err := uc.GetAll("")

		assert.NoError(t, err)
		assert.Equal(t, data, userListDomain)
	})

	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		userRepo.On("GetAll", mock.Anything, mock.Anything).Return([]accounts.Domain{}, errors.New("")).Once()
		data, err := uc.GetAll("")

		assert.Error(t, err)
		assert.Equal(t, data, []accounts.Domain{})
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		setup()
		userRepo.On("GetById", mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		data, err := uc.GetById(1)

		assert.NoError(t, err)
		assert.Equal(t, data, userDomain)
	})

	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		userRepo.On("GetById", mock.Anything, mock.Anything).Return(accounts.Domain{}, errors.New("")).Once()
		data, err := uc.GetById(1)

		assert.Error(t, err)
		assert.Equal(t, data, accounts.Domain{})
	})

	t.Run("Test Case 3", func(t *testing.T) {
		setup()
		mck := accounts.Domain{
			ID: 0,
		}
		userRepo.On("GetById", mock.Anything, mock.Anything).Return(mck, nil).Once()
		data, err := uc.GetById(1)

		assert.Error(t, err)
		assert.Equal(t, data, accounts.Domain{})
	})
}

func TestCreate(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		setup()
		userDomain.Repassword = "z"
		data, err := uc.Create(userDomain)

		assert.Error(t, err)
		assert.Equal(t, data, accounts.Domain{})
	})
	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		userRepo.On("GetByEmail", mock.Anything).Return(accounts.Domain{}, errors.New("")).Once()
		data, err := uc.Create(userDomain)

		assert.Error(t, err)
		assert.Equal(t, data, accounts.Domain{})
	})

	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		userDomain.ID = 1
		userRepo.On("GetByEmail", mock.Anything).Return(userDomain, nil).Once()
		data, err := uc.Create(userDomain)

		assert.Error(t, err)
		assert.Equal(t, data, accounts.Domain{})
	})

	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		cekEmail := accounts.Domain{
			ID: 0,
		}
		userRepo.On("GetByEmail", mock.Anything).Return(cekEmail, nil).Once()
		userRepo.On("Create", mock.Anything).Return(accounts.Domain{}, errors.New("")).Once()
		data, err := uc.Create(userDomain)

		assert.Error(t, err)
		assert.Equal(t, data, accounts.Domain{})
	})

	//t.Run("Test Case 2", func(t *testing.T) {
	//	setup()
	//	cekEmail := accounts.Domain{
	//		ID: 0,
	//	}
	//	userRepo.On("GetByEmail", mock.Anything).Return(cekEmail, nil).Once()
	//	userRepo.On("Create", mock.Anything).Return(userDomain, nil).Once()
	//	data, err := uc.Create(userDomain)
	//
	//	assert.NoError(t, err)
	//	assert.Equal(t, data, userDomain)
	//})

	//t.Run("Test Case 2", func(t *testing.T) {
	//	setup()
	//	userDomain.Password = bcryptPassword
	//	userRepo.On("GetByEmail", mock.Anything).Return(userDomain, nil).Once()
	//	userRepo.On("Create", mock.Anything).Return(accounts.Domain{}, errors.New("")).Once()
	//	data, err := uc.Create(userDomain)
	//
	//	assert.Error(t, err)
	//	assert.Equal(t, data, accounts.Domain{})
	//})

}

func TestUpdate(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		setup()
		userRepo.On("GetById", mock.Anything).Return(accounts.Domain{}, errors.New("")).Once()
		data, err := uc.Update(userDomain)

		assert.Error(t, err)
		assert.Equal(t, data, accounts.Domain{})
	})

	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		userDomain.ID = 0
		userRepo.On("GetById", mock.Anything).Return(userDomain, nil).Once()
		data, err := uc.Update(userDomain)

		assert.Error(t, err)
		assert.Equal(t, data, accounts.Domain{})
	})

	t.Run("Test Case 3", func(t *testing.T) {
		setup()
		userDomain.Name = ""
		userRepo.On("GetById", mock.Anything).Return(userDomain, nil).Once()
		data, err := uc.Update(userDomain)

		assert.Error(t, err)
		assert.Equal(t, data, accounts.Domain{})
	})

	t.Run("Test Case 3", func(t *testing.T) {
		setup()
		userDomain.Handphone = ""
		userRepo.On("GetById", mock.Anything).Return(userDomain, nil).Once()
		data, err := uc.Update(userDomain)

		assert.Error(t, err)
		assert.Equal(t, data, accounts.Domain{})
	})

	t.Run("Test Case 3", func(t *testing.T) {
		setup()
		userDomain.Password = "Syalaala"
		userRepo.On("GetById", mock.Anything).Return(userDomain, nil).Once()
		userRepo.On("UpdateWithPassword", mock.Anything).Return(accounts.Domain{}, errors.New("")).Once()
		data, err := uc.Update(userDomain)

		assert.Error(t, err)
		assert.Equal(t, data, accounts.Domain{})
	})

	t.Run("Test Case 3", func(t *testing.T) {
		setup()
		userRepo.On("GetById", mock.Anything).Return(userDomain, nil).Once()
		userRepo.On("UpdateWithPassword", mock.Anything).Return(userDomain, nil).Once()
		_, err := uc.Update(userDomain)

		assert.Error(t, err)
		//assert.Equal(t, data, userDomain)
	})
}

//}
//
//func TestDelete(t *testing.T) {
//	t.Run("Test Case 1", func(t *testing.T) {
//		setup()
//		userRepo.On("GetById", mock.Anything).Return(accounts.Domain{}, errors.New("")).Once()
//		err := uc.Delete(1)
//
//		assert.Error(t, err)
//	})
//
//	t.Run("Test Case 2", func(t *testing.T) {
//		setup()
//		z := accounts.Domain{
//			ID: 0,
//		}
//		userRepo.On("GetById", mock.Anything).Return(z, nil).Once()
//		err := uc.Delete(1)
//
//		assert.Equal(t, err, baseErrors.ErrNotFound)
//		assert.Error(t, err)
//	})
//
//	t.Run("Test Case 3", func(t *testing.T) {
//		setup()
//		userRepo.On("GetById", mock.Anything).Return(userDomain, nil).Once()
//		userRepo.On("Delete", mock.Anything).Return(errors.New("")).Once()
//
//		err := uc.Delete(1)
//		assert.Error(t, err)
//	})
//
//	t.Run("Test Case 4", func(t *testing.T) {
//		setup()
//		userRepo.On("GetById", mock.Anything).Return(userDomain, nil).Once()
//		userRepo.On("Delete", mock.Anything).Return(nil).Once()
//
//		err := uc.Delete(1)
//		assert.NoError(t, err)
//	})
