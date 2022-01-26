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

func (p *pgOrdersRepo) GetById(id uint) (orders.Domain, error) {
	model := Transactions{}
	e := p.Psql.First(&model, id)
	if e.Error != nil {
		return orders.Domain{}, e.Error
	}
	return model.ToDomain(), nil
}

func (p *pgOrdersRepo) WebhookPaidVA(externalId uint, status string) (orders.Domain, error) {
	model := Transactions{}
	failDate := time.Time{}
	successDate := time.Time{}
	if status == "bayar" {
		successDate = time.Now().Local()
	} else {
		failDate = time.Now().Local()
	}
	e := p.Psql.Model(&model).Where("id = ?", externalId).Updates(Transactions{
		Status:          status,
		SuccessDateTime: successDate,
		FailDateTime:    failDate,
	})

	if e.Error != nil {
		return orders.Domain{}, e.Error
	}
	return model.ToDomain(), nil
}

func (p *pgOrdersRepo) CreateTransaction(domain orders.Domain) (orders.Domain, error) {
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
	err = p.Psql.Preload("DetailTransactions").Preload("Products").Find(&t)
	if err.Error != nil {
		return orders.Domain{}, err.Error
	}
	return t.ToDomain(), nil
}

func (p *pgOrdersRepo) GetHistory(userId uint) ([]orders.Domain, error) {
	var model []Transactions
	e := p.Psql.Preload("Products").Preload("DetailTransactions").Find(&model, "user_id = ?", userId)
	if e.Error != nil {
		return []orders.Domain{}, e.Error
	}
	return ToDomainList(model), nil
}

func (p *pgOrdersRepo) UpdateUniqueValue(transactionId uint, uniqueValue string) error {
	model := DetailTransactions{}
	err := p.Psql.Model(&model).Where("transaction_id = ?", transactionId).Updates(DetailTransactions{
		DetailUniqueValue: uniqueValue,
	})
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (p *pgOrdersRepo) UpdateAfterCreateVA(domain orders.Domain) (orders.Domain, error) {
	t := FromDomainToTransaction(domain)
	model := Transactions{}
	err := p.Psql.Model(&model).Where("id = ?", domain.ID).Updates(Transactions{
		ExpirationPayment: t.ExpirationPayment,
		AccountNumber:     t.AccountNumber,
		BankCode:          t.BankCode,
	})
	if err.Error != nil {
		return orders.Domain{}, err.Error
	}
	return t.ToDomain(), nil
}
