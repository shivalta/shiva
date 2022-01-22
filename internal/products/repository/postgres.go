package repository

import (
	"gorm.io/gorm"
	"shiva/shiva-auth/internal/products"
)

type pgProductsRepo struct {
	Psql *gorm.DB
}

func NewProductsRepo(psql *gorm.DB) products.Repository {
	return &pgProductsRepo{
		Psql: psql,
	}
}

func (p *pgProductsRepo) GetAll(search string, key string) ([]products.Domain, error) {
	var model []Products
	if key == "" {
		err := p.Psql.Find(&model)
		if err.Error != nil {
			return []products.Domain{}, err.Error
		}
		return ToDomainList(model), nil
	}
	err := p.Psql.Find(&model, key+` = ?`, search)
	if err.Error != nil {
		return []products.Domain{}, err.Error
	}
	return ToDomainList(model), nil
}

func (p *pgProductsRepo) GetById(id uint) (products.Domain, error) {
	model := Products{}
	e := p.Psql.First(&model, id)
	if e.Error != nil {
		return products.Domain{}, e.Error
	}
	return model.ToDomain(), nil
}

func (p *pgProductsRepo) Update(d products.Domain) (products.Domain, error) {
	model := Products{}
	e := p.Psql.Model(&model).Where("id = ?", d.ID).Updates(Products{
		ProductClassId:    d.ProductClassId,
		ProductCategoryId: d.ProductCategoryId,
		Sku:               d.Sku,
		Name:              d.Name,
		AdminFee:          d.AdminFee,
		Stock:             d.Stock,
		Price:             d.Price,
		IsActive:          d.IsActive,
	})
	if e.Error != nil {
		return products.Domain{}, e.Error
	}
	return model.ToDomain(), nil
}

func (p *pgProductsRepo) UpdateWithoutImage(d products.Domain) (products.Domain, error) {
	model := Products{}
	e := p.Psql.Model(&model).Where("id = ?", d.ID).Updates(Products{
		ProductClassId:    d.ProductClassId,
		ProductCategoryId: d.ProductCategoryId,
		Sku:               d.Sku,
		Name:              d.Name,
		AdminFee:          d.AdminFee,
		Stock:             d.Stock,
		Price:             d.Price,
		IsActive:          d.IsActive,
	})
	if e.Error != nil {
		return products.Domain{}, e.Error
	}
	return model.ToDomain(), nil
}

func (p *pgProductsRepo) Delete(id uint) error {
	model := Products{}
	e := p.Psql.Delete(&model, id)
	if e.Error != nil {
		return e.Error
	}
	return nil
}

func (p *pgProductsRepo) Create(d products.Domain) (products.Domain, error) {
	u := FromDomain(d)
	err := p.Psql.Create(&u)
	if err.Error != nil {
		return products.Domain{}, err.Error
	}
	return u.ToDomain(), nil
}
