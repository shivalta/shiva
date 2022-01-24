package usecase

import (
	"github.com/go-playground/validator/v10"
	"shiva/shiva-auth/internal/orders"
)

type Usecase struct {
	data     orders.Repository
	xendit   orders.XenditRepository
	validate *validator.Validate
}

func NewOrderUsecase(data orders.Repository, xendit orders.XenditRepository) orders.Usecase {
	return &Usecase{
		data:     data,
		xendit:   xendit,
		validate: validator.New(),
	}
}

func (u Usecase) CheckoutPulsa(userId uint, productId uint) ([]orders.Domain, error) {
	panic("implement me")
}

func (u Usecase) CheckoutPDAM(userId uint, productId uint) ([]orders.Domain, error) {
	panic("implement me")
}

func (u Usecase) CheckoutListrik(userId uint, productId uint) ([]orders.Domain, error) {
	panic("implement me")
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
