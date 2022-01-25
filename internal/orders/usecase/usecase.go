package usecase

import (
	"github.com/go-playground/validator/v10"
	"shiva/shiva-auth/internal/orders"
	"shiva/shiva-auth/internal/products"
	"shiva/shiva-auth/utils/baseErrors"
	"strings"
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
		order.UserValue = userValue + ` - ` + prod.ProductCategory.Name
	} else if prodCat == "token" {
		order.TotalPrice = (*prod.Price + prod.AdminFee) + int(float32(*prod.Price)*(prod.ProductCategory.Tax/float32(100)))
		m, err := u.mockapi.GetMockListrik(userValue)
		if err != nil {
			return orders.Domain{}, err
		}
		order.UserValue = userValue + ` - ` + m.UserValue
	} else {
		m, err := u.mockapi.GetMockPDAM(userValue)
		if err != nil {
			return orders.Domain{}, err
		}
		order.UserValue = userValue + ` - ` + m.UserValue
		order.TotalPrice = (m.TotalPrice + prod.AdminFee) + int(prod.ProductCategory.Tax/float32(100))
	}
	order.TotalTax = prod.ProductCategory.Tax
	order.TotalAdmin = prod.AdminFee
	return order, nil
}

func (u Usecase) PaymentChannels() ([]orders.Domain, error) {
	paymentMethod, err := u.xendit.PaymentChannels()
	if err != nil {
		return []orders.Domain{}, err
	}
	return paymentMethod, nil
}

func (u Usecase) CreateVA(productId uint, userId uint, bankCode string, userValue string) (orders.Domain, error) {
	//res, err := u.data.CreateTransaction(productId, userId, bankCode)
	//if err != nil {
	//	return orders.Domain{}, err
	//}
	//		AccountNumber:     p.AccountNumber,
	//		BankName:          p.Name,
	//		ExpirationPayment: p.ExpirationDate,
	//va := u.xendit.CreateVA()
	panic("")
}

func (u Usecase) WebhookCreateVA(domain orders.Domain) (orders.Domain, error) {
	panic("implement me")
}

func (u Usecase) WebhookPaidVA(domain orders.Domain) (orders.Domain, error) {
	panic("implement me")
}
