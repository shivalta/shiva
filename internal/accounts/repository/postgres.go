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

func (p *pgAccountRepository) GetByEmail(email string) (accounts.Domain, error) {
	model := Users{}
	e := p.Psql.First(&model, "email = ?", email)
	if e.Error != nil{
		return accounts.Domain{}, e.Error
	}
	return model.UserToDomain(), nil
}

func (p *pgAccountRepository) ChangePassword(id uint, password string) (accounts.Domain, error) {
	model := Users{}
	e := p.Psql.Model(&model).Where("id = ?", id).Updates(Users{
		Password: password,
	})
	if e.Error != nil{
		return accounts.Domain{}, e.Error
	}
	return model.UserToDomain(), nil
}

func (p *pgAccountRepository) GetById(id uint) (accounts.Domain, error) {
	model := Users{}
	e := p.Psql.First(&model, id)
	if e.Error != nil{
		return accounts.Domain{}, e.Error
	}
	return model.UserToDomain(), nil
}

func (p *pgAccountRepository) Update(user accounts.Domain) (accounts.Domain, error) {
	model := Users{}
	//p.Psql.First(&model, user.ID)
	//model.Name = user.Name
	//model.Address = user.Address
	//model.Handphone = user.Handphone
	//p.Psql.Save(&model)
	e := p.Psql.Model(&model).Where("id = ?", user.ID).Updates(Users{
		Name:      user.Name,
		Handphone: user.Handphone,
		Address:   user.Address,
	})
	if e.Error != nil{
		return accounts.Domain{}, e.Error
	}
	return model.UserToDomain(), nil
}

func (p *pgAccountRepository) Delete(id uint) error {
	model := Users{}
	e := p.Psql.Delete(&model, id)
	if e.Error != nil {
		return e.Error
	}
	return nil
}

func (p *pgAccountRepository) Create(user accounts.Domain) (accounts.Domain, error) {
	u := FromDomain(user)
	err := p.Psql.Create(&u)
	if err.Error != nil {
		return accounts.Domain{}, err.Error
	}
	return u.UserToDomain(), nil
}

func (p *pgAccountRepository) GetAll(search string) ([]accounts.Domain, error) {
	var user []Users
	err := p.Psql.Find(&user)

	if err.Error != nil {
		return []accounts.Domain{}, err.Error
	}
	return UserToDomainList(user), nil
}
