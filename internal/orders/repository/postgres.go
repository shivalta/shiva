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

func (p pgOrdersRepo) CreateTransaction(userId uint, bankCode string, domain orders.Domain) (orders.Domain, error) {
	panic("implement me")
}

//func (p pgOrdersRepo) CreateTransaction(productId uint, userId uint, bankCode string) (orders.Domain, error) {
//	u := FromDomain(user)
//	err := p.Psql.Create(&u)
//	if err.Error != nil {
//		return accounts.Domain{}, err.Error
//	}
//	return u.UserToDomain(), nil
//	panic("")
//
//}

func (p pgOrdersRepo) WebhookCreateVA(domain orders.Domain) (orders.Domain, error) {
	panic("implement me")
}

func (p pgOrdersRepo) WebhookPaidVA(domain orders.Domain) (orders.Domain, error) {
	panic("implement me")
}
