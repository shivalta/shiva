package repository

import (
	"gorm.io/gorm"
	"shiva/shiva-auth/internal/orders"
)

type pgOrdersRepo struct {
	Psql *gorm.DB
}

func NewOrdersRepo(psql *gorm.DB) orders.Repository {
	return &pgOrdersRepo{
		Psql: psql,
	}
}

func (p pgOrdersRepo) CheckoutPulsa(userId uint, productId uint) (orders.Domain, error) {
	panic("implement me")
}

func (p pgOrdersRepo) CheckoutPDAM(userId uint, productId uint) (orders.Domain, error) {
	panic("implement me")
}

func (p pgOrdersRepo) CheckoutListrik(userId uint, productId uint) (orders.Domain, error) {
	panic("implement me")
}

func (p pgOrdersRepo) CreateTransaction(productId uint, userId uint, bankCode string) (orders.Domain, error) {
	panic("implement me")
}

func (p pgOrdersRepo) WebhookCreateVA(domain orders.Domain) (orders.Domain, error) {
	panic("implement me")
}

func (p pgOrdersRepo) WebhookPaidVA(domain orders.Domain) (orders.Domain, error) {
	panic("implement me")
}
