package repository

import (
	"gorm.io/gorm"
	"shiva/shiva-auth/internal/categories"
)

type pgCategoriesRepo struct {
	Psql *gorm.DB
}

func NewCategoriesRepo(psql *gorm.DB) categories.Repository {
	return &pgCategoriesRepo{
		Psql: psql,
	}
}

func (p *pgCategoriesRepo) GetAll(search string, key string) ([]categories.Domain, error) {
	var user []ProductCategories
	err := p.Psql.Find(&user)

	if err.Error != nil {
		return []categories.Domain{}, err.Error
	}
	return ToDomainList(user), nil
}

func (p *pgCategoriesRepo) GetById(id uint) (categories.Domain, error) {
	model := ProductCategories{}
	e := p.Psql.First(&model, id)
	if e.Error != nil {
		return categories.Domain{}, e.Error
	}
	return model.ToDomain(), nil
}

func (p *pgCategoriesRepo) Update(d categories.Domain) (categories.Domain, error) {
	model := ProductCategories{}
	e := p.Psql.Model(&model).Where("id = ?", d.ID).Updates(ProductCategories{
		ProductClassId: d.ProductClassId,
		Name:           d.Name,
		Image:          d.ImageUrl,
		Tax:            d.Tax,
	})
	if e.Error != nil {
		return categories.Domain{}, e.Error
	}
	return model.ToDomain(), nil
}

func (p *pgCategoriesRepo) UpdateWithoutImage(d categories.Domain) (categories.Domain, error) {
	model := ProductCategories{}
	e := p.Psql.Model(&model).Where("id = ?", d.ID).Updates(ProductCategories{
		ProductClassId: d.ProductClassId,
		Name:           d.Name,
		Tax:            d.Tax,
	})
	if e.Error != nil {
		return categories.Domain{}, e.Error
	}
	return model.ToDomain(), nil
}

func (p *pgCategoriesRepo) Delete(id uint) error {
	model := ProductCategories{}
	e := p.Psql.Delete(&model, id)
	if e.Error != nil {
		return e.Error
	}
	return nil
}

func (p *pgCategoriesRepo) Create(d categories.Domain) (categories.Domain, error) {
	u := FromDomain(d)
	err := p.Psql.Create(&u)
	if err.Error != nil {
		return categories.Domain{}, err.Error
	}
	return u.ToDomain(), nil
}
