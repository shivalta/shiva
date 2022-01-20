package repository

import (
	"gorm.io/gorm"
	"shiva/shiva-auth/internal/class"
)

type pgClassRepo struct {
	Psql *gorm.DB
}

func NewClassRepo(psql *gorm.DB) class.Repository {
	return &pgClassRepo{
		Psql: psql,
	}
}

func (p *pgClassRepo) GetAll(search string, key string) ([]class.Domain, error) {
	var user []ProductClass
	err := p.Psql.Find(&user)

	if err.Error != nil {
		return []class.Domain{}, err.Error
	}
	return ToDomainList(user), nil
}

func (p *pgClassRepo) GetById(id uint) (class.Domain, error) {
	model := ProductClass{}
	e := p.Psql.First(&model, id)
	if e.Error != nil {
		return class.Domain{}, e.Error
	}
	return model.ToDomain(), nil
}

func (p *pgClassRepo) Update(d class.Domain) (class.Domain, error) {
	model := ProductClass{}
	e := p.Psql.Model(&model).Where("id = ?", d.ID).Updates(ProductClass{
		Name:    d.Name,
		IsPasca: d.IsPasca,
		Image:   d.Image,
	})
	if e.Error != nil {
		return class.Domain{}, e.Error
	}
	return model.ToDomain(), nil
}

func (p *pgClassRepo) Delete(id uint) error {
	model := ProductClass{}
	e := p.Psql.Delete(&model, id)
	if e.Error != nil {
		return e.Error
	}
	return nil
}

func (p *pgClassRepo) Create(d class.Domain) (class.Domain, error) {
	u := FromDomain(d)
	err := p.Psql.Create(&u)
	if err.Error != nil {
		return class.Domain{}, err.Error
	}
	return u.ToDomain(), nil
}
