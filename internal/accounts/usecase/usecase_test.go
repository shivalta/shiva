package usecase_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"shiva/shiva-auth/cmd/http/middlewares"
	"shiva/shiva-auth/internal/accounts"
	"shiva/shiva-auth/internal/accounts/mocks"
	"shiva/shiva-auth/internal/accounts/usecase"
	"shiva/shiva-auth/utils/baseErrors"
	"testing"
)

var uc accounts.Usecase
var userRepo mocks.Repository
var userDomain accounts.Domain
var userListDomain []accounts.Domain
var bcryptPassword string
var base64Email string
var encryptEmail string
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
		IsActive:   true,
		Password:   "123",
		Repassword: "123",
	}
	base64Email = "ZHdpa3kuZGV2QGdtYWlsLmNvbQ=="
	encryptEmail = "b846a0b52a9f6c8184b3f7877fb33889dd055c306ed0d3c1c2dd2831881a4ee902898a51e5bb4dbcdda9d82cb62908c64e028f98182521494805a161552b4ff255617d4ced09f2f89c5d58f649cec7c21e988c13"
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
		data, err := uc.Update(userDomain)

		assert.NoError(t, err)
		assert.Equal(t, data, userDomain)
	})

	t.Run("Test Case 3", func(t *testing.T) {
		setup()
		userDomain.Password = "ZZZ"
		userRepo.On("GetById", mock.Anything).Return(userDomain, nil).Once()
		data, err := uc.Update(userDomain)

		assert.Error(t, err)
		assert.Equal(t, data, accounts.Domain{})
	})

	t.Run("Test Case 3", func(t *testing.T) {
		setup()
		userDomain.Password = ""
		userRepo.On("GetById", mock.Anything).Return(userDomain, nil).Once()
		userRepo.On("Update", mock.Anything).Return(accounts.Domain{}, errors.New("")).Once()
		data, err := uc.Update(userDomain)

		assert.Error(t, err)
		assert.Equal(t, data, accounts.Domain{})
	})

	t.Run("Test Case 3", func(t *testing.T) {
		setup()
		userDomain.Password = ""
		userRepo.On("GetById", mock.Anything).Return(userDomain, nil).Once()
		userRepo.On("Update", mock.Anything).Return(userDomain, nil).Once()
		data, err := uc.Update(userDomain)

		assert.NoError(t, err)
		assert.Equal(t, data, userDomain)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		setup()
		userRepo.On("GetById", mock.Anything).Return(accounts.Domain{}, errors.New("")).Once()
		err := uc.Delete(1)

		assert.Error(t, err)
	})

	t.Run("Test Case 2", func(t *testing.T) {
		setup()
		z := accounts.Domain{
			ID: 0,
		}
		userRepo.On("GetById", mock.Anything).Return(z, nil).Once()
		err := uc.Delete(1)

		assert.Equal(t, err, baseErrors.ErrNotFound)
		assert.Error(t, err)
	})

	t.Run("Test Case 3", func(t *testing.T) {
		setup()
		userRepo.On("GetById", mock.Anything).Return(userDomain, nil).Once()
		userRepo.On("Delete", mock.Anything).Return(errors.New("")).Once()

		err := uc.Delete(1)
		assert.Error(t, err)
	})

	t.Run("Test Case 4", func(t *testing.T) {
		setup()
		userRepo.On("GetById", mock.Anything).Return(userDomain, nil).Once()
		userRepo.On("Delete", mock.Anything).Return(nil).Once()

		err := uc.Delete(1)
		assert.NoError(t, err)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Test Case", func(t *testing.T) {
		setup()
		userDomain.Password = ""
		data, tok, err := uc.Login(userDomain.Email, userDomain.Password)

		assert.Error(t, err)
		assert.Equal(t, data, accounts.Domain{})
		assert.Equal(t, tok, "")
	})

	t.Run("Test Case", func(t *testing.T) {
		setup()
		userDomain.Email = ""
		userDomain.Password = "Salalla"
		data, tok, err := uc.Login(userDomain.Email, userDomain.Password)

		assert.Error(t, err)
		assert.Equal(t, data, accounts.Domain{})
		assert.Equal(t, tok, "")
	})

	t.Run("Test Case", func(t *testing.T) {
		setup()
		userRepo.On("GetByEmail", mock.Anything).Return(accounts.Domain{}, nil).Once()
		data, tok, err := uc.Login("shiva@shiva.co", "syalalahaik")

		assert.Error(t, err)
		assert.Equal(t, data, accounts.Domain{})
		assert.Equal(t, tok, "")
	})

	t.Run("Test Case", func(t *testing.T) {
		setup()
		userDomain.IsActive = false
		userRepo.On("GetByEmail", mock.Anything).Return(userDomain, nil).Once()
		data, tok, err := uc.Login("shiva@shiva.co", "syalalahaik")

		assert.Error(t, err)
		assert.Equal(t, data, accounts.Domain{})
		assert.Equal(t, tok, "")
		assert.Equal(t, err, baseErrors.ErrUserNotActive)
	})
	t.Run("Test Case", func(t *testing.T) {
		setup()
		userDomain.IsActive = true
		userRepo.On("GetByEmail", mock.Anything).Return(userDomain, nil).Once()
		_, tok, _ := uc.Login(userDomain.Email, userDomain.Password)

		assert.NotNil(t, tok, token)
	})

	t.Run("Test Case", func(t *testing.T) {
		setup()
		userDomain.IsActive = true
		userDomain.Password = bcryptPassword
		userRepo.On("GetByEmail", mock.Anything).Return(userDomain, nil).Once()
		data, tok, err := uc.Login(userDomain.Email, "123")

		assert.NoError(t, err)
		assert.NotNil(t, tok, token)
		assert.Equal(t, userDomain, data)
	})
}

func TestVerify(t *testing.T) {
	t.Run("Test Case", func(t *testing.T) {
		setup()
		_, err := uc.Verify(base64Email, encryptEmail)
		assert.Error(t, err)
	})
}
