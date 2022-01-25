package repository

import (
	"gorm.io/gorm"
	"shiva/shiva-auth/internal/orders"
	"time"
)

type pgOrdersRepo struct {
	Psql *gorm.DB
}

func NewOrdersRepo(psql *gorm.DB) orders.Repository {
	return &pgOrdersRepo{
		Psql: psql,
	}
}

func (p pgOrdersRepo) CreateTransaction(domain orders.Domain) (orders.Domain, error) {
	t := FromDomainToTransaction(domain)
	d := FromDomainToDetail(domain)
	t.PendingDateTime = time.Now().Local()
	t.Status = "pending"
	err := p.Psql.Create(&t)
	if err.Error != nil {
		return orders.Domain{}, err.Error
	}
	d.TransactionId = t.ID
	err = p.Psql.Create(&d)
	if err.Error != nil {
		return orders.Domain{}, err.Error
	}
	return t.ToDomain(), nil
}

func (p pgOrdersRepo) UpdateAfterCreateVA(domain orders.Domain) (orders.Domain, error) {
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
