package repository

import (
	"gorm.io/gorm"
	"shiva/shiva-auth/internal/accounts"
)

type pgAccountRepository struct {
	Psql *gorm.DB
}

func NewAccountRepo(psql *gorm.DB) accounts.Repository {
	return &pgAccountRepository{
		Psql: psql,
	}
}

func (p *pgAccountRepository) Create(user accounts.Domain) (accounts.Domain, error) {
	panic("implement me")
}

func (p *pgAccountRepository) GetAll(search string) ([]accounts.Domain, error) {
	var user []Users
	err := p.Psql.Find(&user)

	if err.Error != nil {
		return []accounts.Domain{}, err.Error
	}
	return UserToDomainList(user), nil
}
