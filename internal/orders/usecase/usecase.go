package usecase

import (
	"github.com/go-playground/validator/v10"
	"shiva/shiva-auth/internal/orders"
	"shiva/shiva-auth/internal/products"
	"shiva/shiva-auth/utils/baseErrors"
)

type Usecase struct {
	data     orders.Repository
	xendit   orders.XenditRepository
	mockapi  orders.MockupIoRepository
	product  products.Usecase
	validate *validator.Validate
}

func NewOrderUsecase(data orders.Repository, xendit orders.XenditRepository, mockapi orders.MockupIoRepository, product products.Usecase) orders.Usecase {
	return &Usecase{
		data:     data,
		xendit:   xendit,
		mockapi:  mockapi,
		product:  product,
		validate: validator.New(),
	}
}

func (u Usecase) CheckoutPulsa(userValue string, productId uint) (orders.Domain, error) {
	if userValue == "" {
		return orders.Domain{}, baseErrors.ErrNoHpRequired
	}
	prod, err := u.product.GetById(productId)
	if err != nil {
		return orders.Domain{}, err
	}
	totalPrice := (prod.Price + prod.AdminFee) * int(prod.ProductCategory.Tax/float32(100))
	order := orders.Domain{}
	order.TotalPrice = totalPrice
	order.Products = orders.Products{
		ID: prod.ID,
		ProductClass: orders.Class{
			ID:       prod.ProductClass.ID,
			Name:     prod.ProductClass.Name,
			IsPasca:  prod.ProductClass.IsPasca,
			ImageUrl: prod.ProductClass.ImageUrl,
		},
		ProductCategory: orders.Categories{
			ID:             prod.ProductCategory.ID,
			ProductClassId: prod.ProductCategory.ProductClassId,
			Name:           prod.ProductCategory.Name,
			ImageUrl:       prod.ProductCategory.ImageUrl,
			Tax:            prod.ProductCategory.Tax,
		},
		Sku:      prod.Sku,
		Name:     prod.Name,
		AdminFee: prod.AdminFee,
		Stock:    prod.Stock,
		Price:    prod.Price,
		IsActive: prod.IsActive,
	}
	order.UserValue = userValue
	return order, nil
}

func (u Usecase) CheckoutListrik(userValue string, productId uint) (orders.Domain, error) {
	if userValue == "" {
		return orders.Domain{}, baseErrors.ErrNoHpRequired
	}
	prod, err := u.product.GetById(productId)
	if err != nil {
		return orders.Domain{}, err
	}
	totalPrice := (prod.Price + prod.AdminFee) * int(prod.ProductCategory.Tax/float32(100))
	order := orders.Domain{}
	order.TotalPrice = totalPrice
	order.Products = orders.Products{
		ID: prod.ID,
		ProductClass: orders.Class{
			ID:       prod.ProductClass.ID,
			Name:     prod.ProductClass.Name,
			IsPasca:  prod.ProductClass.IsPasca,
			ImageUrl: prod.ProductClass.ImageUrl,
		},
		ProductCategory: orders.Categories{
			ID:             prod.ProductCategory.ID,
			ProductClassId: prod.ProductCategory.ProductClassId,
			Name:           prod.ProductCategory.Name,
			ImageUrl:       prod.ProductCategory.ImageUrl,
			Tax:            prod.ProductCategory.Tax,
		},
		Sku:      prod.Sku,
		Name:     prod.Name,
		AdminFee: prod.AdminFee,
		Stock:    prod.Stock,
		Price:    prod.Price,
		IsActive: prod.IsActive,
	}
	order.UserValue = userValue
	return order, nil
}

func (u Usecase) CheckoutPDAM(userValue string, productId uint) (orders.Domain, error) {
	if userValue == "" {
		return orders.Domain{}, baseErrors.ErrNoHpRequired
	}
	prod, err := u.product.GetById(productId)
	if err != nil {
		return orders.Domain{}, err
	}
	totalPrice := (prod.Price + prod.AdminFee) * int(prod.ProductCategory.Tax/float32(100))
	order := orders.Domain{}
	order.TotalPrice = totalPrice
	order.Products = orders.Products{
		ID: prod.ID,
		ProductClass: orders.Class{
			ID:       prod.ProductClass.ID,
			Name:     prod.ProductClass.Name,
			IsPasca:  prod.ProductClass.IsPasca,
			ImageUrl: prod.ProductClass.ImageUrl,
		},
		ProductCategory: orders.Categories{
			ID:             prod.ProductCategory.ID,
			ProductClassId: prod.ProductCategory.ProductClassId,
			Name:           prod.ProductCategory.Name,
			ImageUrl:       prod.ProductCategory.ImageUrl,
			Tax:            prod.ProductCategory.Tax,
		},
		Sku:      prod.Sku,
		Name:     prod.Name,
		AdminFee: prod.AdminFee,
		Stock:    prod.Stock,
		Price:    prod.Price,
		IsActive: prod.IsActive,
	}
	pdam, err := u.mockapi.GetName(userValue)
	if err != nil {
		return orders.Domain{}, err
	}
	order.UserValue = userValue + ` - ` + pdam.UserValue
	return order, nil
}

func (u Usecase) PaymentChannel() (orders.Domain, error) {
	panic("implement me")
}

func (u Usecase) CreateVA(productId uint, userId uint, bankCode string) (orders.Domain, error) {
	panic("implement me")
}

func (u Usecase) WebhookCreateVA(domain orders.Domain) (orders.Domain, error) {
	panic("implement me")
}

func (u Usecase) WebhookPaidVA(domain orders.Domain) (orders.Domain, error) {
	panic("implement me")
}
