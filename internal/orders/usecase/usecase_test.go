package usecase_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"shiva/shiva-auth/internal/accounts"
	usermock "shiva/shiva-auth/internal/accounts/mocks"
	"shiva/shiva-auth/internal/orders"
	"shiva/shiva-auth/internal/orders/mocks"
	"shiva/shiva-auth/internal/orders/usecase"
	"shiva/shiva-auth/internal/products"
	prodmock "shiva/shiva-auth/internal/products/mocks"
	"testing"
)

var uc orders.Usecase
var usecaseUser usermock.Usecase
var usecaseProduct prodmock.Usecase
var orderRepo mocks.Repository
var xenditRepo mocks.XenditRepository
var mockapiRepo mocks.MockupIoRepository
var orderDomain orders.Domain
var orderListDomain []orders.Domain
var userDomain accounts.Domain
var productDomain products.Domain

func setup() {
	uc = usecase.NewOrdersUsecase(&usecaseUser, &orderRepo, &xenditRepo, &mockapiRepo, &usecaseProduct)
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
		Price:             nil,
		IsActive:          false,
	}
	orderDomain = orders.Domain{
		ID:            1,
		Status:        "bayar",
		UserId:        1,
		TotalPrice:    1,
		TotalTax:      1,
		TotalAdmin:    1,
		AccountNumber: "TEST",
		UserValue:     "123",
		UniqueValue:   "123",
		BankName:      "MANDIRI",
		BankCode:      "MANDIRI",
		Amount:        10000,
	}
	orderListDomain = append(orderListDomain, orderDomain)
}

func TestCheckout(t *testing.T) {
	t.Run("Test Case", func(t *testing.T) {
		setup()
		data, err := uc.Checkout("", 0)
		if err != nil {
			return
		}

		assert.Error(t, err)
		assert.Equal(t, data, orders.Domain{})
	})
}

func TestPaymentChannels(t *testing.T) {
	t.Run("Test Case", func(t *testing.T) {
		setup()
		xenditRepo.On("PaymentChannels").Return([]orders.Domain{}, errors.New("")).Once()
		data, err := uc.PaymentChannels()

		assert.Error(t, err)
		assert.Equal(t, data, []orders.Domain{})
	})

	t.Run("Test Case", func(t *testing.T) {
		setup()
		xenditRepo.On("PaymentChannels").Return(orderListDomain, nil).Once()
		data, err := uc.PaymentChannels()

		assert.NoError(t, err)
		assert.Equal(t, data, orderListDomain)
	})
}

func TestGetHistory(t *testing.T) {
	t.Run("Test Case", func(t *testing.T) {
		setup()
		orderRepo.On("GetHistory", mock.Anything).Return([]orders.Domain{}, errors.New("")).Once()
		data, err := uc.GetHistory(userDomain.ID)

		assert.Error(t, err)
		assert.Equal(t, data, []orders.Domain{})
	})

	t.Run("Test Case", func(t *testing.T) {
		setup()
		orderRepo.On("GetHistory", mock.Anything).Return(orderListDomain, nil).Once()
		data, err := uc.GetHistory(userDomain.ID)

		assert.NoError(t, err)
		assert.Equal(t, data, orderListDomain)
	})
}

//t.Run("Test Case", func(t *testing.T) {
//	setup()
//	productDomain.ProductClass.Name = "token"
//	usecaseProduct.On("GetById", mock.Anything).Return(productDomain, nil).Once()
//	mockapiRepo.On("GetMockListrik", mock.Anything).Return(orders.Domain{}, errors.New("")).Once()
//	data, err := uc.Checkout(orderDomain.UserValue, 1)
//
//	assert.Error(t, err)
//	assert.Equal(t, data, orders.Domain{})
//})

//t.Run("Test Case 2", func(t *testing.T) {
//	setup()
//	orderRepo.On("GetAll", mock.Anything, mock.Anything).Return([]orders.Domain{}, errors.New("")).Once()
//	data, err := uc.GetAll("", "")
//
//	assert.Error(t, err)
//	assert.Equal(t, data, []orders.Domain{})
//})
//}
//	CreateVA(orderId uint, userId uint, bankCode string, userValue string) (Domain, error)
//	funPaymentChannels() ([]Domain, error)
//	GetHistory(userId uint) ([]Domain, error)
//	WebhookPaidVA(externalId uint, amount int) (string, error)
