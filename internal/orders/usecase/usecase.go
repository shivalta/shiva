package usecase

import (
	"github.com/go-playground/validator/v10"
	"shiva/shiva-auth/internal/orders"
	"shiva/shiva-auth/internal/products"
	"shiva/shiva-auth/utils/baseErrors"
	"shiva/shiva-auth/utils/generator"
	"strconv"
	"strings"
	"time"
)

type Usecase struct {
	data     orders.Repository
	xendit   orders.XenditRepository
	mockapi  orders.MockupIoRepository
	product  products.Usecase
	validate *validator.Validate
}

func NewOrdersUsecase(data orders.Repository, xendit orders.XenditRepository, mockapi orders.MockupIoRepository, product products.Usecase) orders.Usecase {
	return &Usecase{
		data:     data,
		xendit:   xendit,
		mockapi:  mockapi,
		product:  product,
		validate: validator.New(),
	}
}

func (u Usecase) Checkout(userValue string, productId uint) (orders.Domain, error) {
	if userValue == "" {
		return orders.Domain{}, baseErrors.ErrNoHpRequired
	}
	prod, err := u.product.GetById(productId)
	if err != nil {
		return orders.Domain{}, err
	}
	order := orders.Domain{}
	prodCat := strings.ToLower(prod.ProductClass.Name)
	if prodCat == "pulsa" {
		order.TotalPrice = (*prod.Price + prod.AdminFee) + int(float32(*prod.Price)*(prod.ProductCategory.Tax/float32(100)))
		order.UserValue = prod.ProductCategory.Name
	} else if prodCat == "token" {
		order.TotalPrice = (*prod.Price + prod.AdminFee) + int(float32(*prod.Price)*(prod.ProductCategory.Tax/float32(100)))
		m, err := u.mockapi.GetMockListrik(userValue)
		if err != nil {
			return orders.Domain{}, err
		}
		order.UserValue = m.UserValue
	} else {
		m, err := u.mockapi.GetMockPDAM(userValue)
		if err != nil {
			return orders.Domain{}, err
		}
		order.UserValue = m.UserValue
		order.TotalPrice = (m.TotalPrice + prod.AdminFee) + int(prod.ProductCategory.Tax/float32(100))
	}
	order.TotalTax = prod.ProductCategory.Tax
	order.TotalAdmin = prod.AdminFee
	return order, nil
}

func (u Usecase) GetHistory(userId uint) ([]orders.Domain, error) {
	h, err := u.data.GetHistory(userId)
	if err != nil {
		return []orders.Domain{}, err
	}
	return h, nil
}

func (u Usecase) PaymentChannels() ([]orders.Domain, error) {
	paymentMethod, err := u.xendit.PaymentChannels()
	if err != nil {
		return []orders.Domain{}, err
	}
	return paymentMethod, nil
}

func (u Usecase) CreateVA(productId uint, userId uint, bankCode string, userValue string) (orders.Domain, error) {
	if userValue == "" {
		return orders.Domain{}, baseErrors.ErrNoHpRequired
	}
	prod, err := u.product.GetById(productId)
	if err != nil {
		return orders.Domain{}, err
	}

	order := orders.Domain{
		Products: orders.Products{
			ID: prod.ID,
			ProductClass: orders.Class{
				ID:       prod.ProductClass.ID,
				Name:     prod.ProductClass.Name,
				IsPasca:  prod.ProductClass.IsPasca,
				ImageUrl: prod.ProductClass.ImageUrl,
			},
			ProductClassId: prod.ProductClassId,
			ProductCategory: orders.Categories{
				ID:             prod.ProductCategory.ID,
				ProductClassId: prod.ProductCategory.ProductClassId,
				Name:           prod.ProductCategory.Name,
				ImageUrl:       prod.ProductCategory.ImageUrl,
				Tax:            prod.ProductCategory.Tax,
			},
			ProductCategoryId: prod.ProductCategoryId,
			Sku:               prod.Sku,
			Name:              prod.Name,
			AdminFee:          prod.AdminFee,
			Stock:             prod.Stock,
			Price:             prod.Price,
			IsActive:          prod.IsActive,
		},
		UserId:     userId,
		BankCode:   bankCode,
		IsLoggedin: false,
	}

	prodCat := strings.ToLower(prod.ProductClass.Name)

	if prodCat == "pulsa" {
		order.TotalPrice = (*prod.Price + prod.AdminFee) + int(float32(*prod.Price)*(prod.ProductCategory.Tax/float32(100)))
		order.UserValue = prod.ProductCategory.Name
	} else if prodCat == "token" {
		order.TotalPrice = (*prod.Price + prod.AdminFee) + int(float32(*prod.Price)*(prod.ProductCategory.Tax/float32(100)))
		m, err := u.mockapi.GetMockListrik(userValue)
		if err != nil {
			return orders.Domain{}, err
		}
		order.UserValue = m.UserValue
	} else {
		m, err := u.mockapi.GetMockPDAM(userValue)
		if err != nil {
			return orders.Domain{}, err
		}
		order.UserValue = m.UserValue
		order.TotalPrice = (m.TotalPrice + prod.AdminFee) + int(prod.ProductCategory.Tax/float32(100))
	}

	res, err := u.data.CreateTransaction(order)
	if err != nil {
		return orders.Domain{}, err
	}

	xendit, err := u.xendit.CreateVA(strconv.Itoa(int(res.ID)), bankCode)

	order.ID = res.ID
	order.BankName = xendit.BankName
	order.AccountNumber = xendit.AccountNumber
	order.ExpirationPayment = xendit.ExpirationPayment
	order.Status = res.Status

	_, err = u.data.UpdateAfterCreateVA(order)
	if err != nil {
		return orders.Domain{}, err
	}

	return order, nil
}

func (u Usecase) WebhookPaidVA(externalId uint, amount int) (string, error) {
	t, err := u.data.GetById(externalId)
	status := ""
	if err != nil {
		return "", err
	}
	if t.ExpirationPayment.Unix() < time.Now().Local().Unix() {
		return "", baseErrors.ErrExpiredPay
	}
	if t.TotalPrice != amount {
		status = "kadaluarsa"
		_, err := u.data.WebhookPaidVA(externalId, "kadaluarsa")
		if err != nil {
			return "", err
		}
		return "", baseErrors.ErrAmountNotMatch
	} else {
		status = "bayar"
	}
	_, err = u.data.WebhookPaidVA(externalId, "bayar")
	if err != nil {
		return "", err
	}

	if t.Products.ProductClass.Name == "pln" {
		token, err := generator.GenerateToken()
		if err != nil {
			return "", err
		}
		err = u.data.UpdateUniqueValue(externalId, token)
		if err != nil {
			return "", err
		}
	}
	return status, nil
}
