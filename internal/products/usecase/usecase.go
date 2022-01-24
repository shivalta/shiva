package Usecase

import (
	"github.com/go-playground/validator/v10"
	"shiva/shiva-auth/internal/categories"
	"shiva/shiva-auth/internal/class"
	"shiva/shiva-auth/internal/products"
	"shiva/shiva-auth/utils/baseErrors"
)

type Usecase struct {
	data            products.Repository
	validate        *validator.Validate
	classUsecase    class.Usecase
	categoryUsecase categories.Usecase
}

func NewProductsUsecase(r products.Repository, classUsecase class.Usecase, categoryUsecase categories.Usecase) products.Usecase {
	return &Usecase{
		data:            r,
		validate:        validator.New(),
		classUsecase:    classUsecase,
		categoryUsecase: categoryUsecase,
	}
}

func (uc Usecase) GetAll(search string, key string) (data []products.Domain, err error) {
	res, err := uc.data.GetAll(search, key)
	if err != nil {
		return []products.Domain{}, err
	}
	return res, nil
}

func (uc Usecase) GetById(id uint) (products.Domain, error) {
	u, err := uc.data.GetById(id)
	if err != nil {
		return products.Domain{}, err
	} else if u.ID == 0 {
		return products.Domain{}, baseErrors.ErrNotFound
	}
	return u, nil
}

func (uc Usecase) Create(d products.Domain) (products.Domain, error) {
	cls, err := uc.classUsecase.GetById(d.ProductClassId)
	if err != nil {
		return products.Domain{}, err
	}
	if cls.IsPasca {
		d.Price = nil
	}
	category, err := uc.categoryUsecase.GetById(d.ProductCategoryId)
	if err != nil {
		return products.Domain{}, err
	}
	p, err := uc.data.Create(d)
	if err != nil {
		return products.Domain{}, err
	}
	p.ProductClass = products.Class{
		ID:       cls.ID,
		Name:     cls.Name,
		IsPasca:  cls.IsPasca,
		ImageUrl: cls.ImageUrl,
		Slug:     cls.Slug,
	}
	p.ProductCategory = products.Categories{
		ID:             category.ID,
		ProductClassId: category.ProductClassId,
		Name:           category.Name,
		ImageUrl:       category.ImageUrl,
		Slug:           category.Slug,
		Tax:            category.Tax,
	}
	return p, nil
}

func (uc Usecase) Update(d products.Domain) (products.Domain, error) {
	_, err := uc.data.GetById(d.ID)
	data, err := uc.data.UpdateWithoutImage(d)
	if err != nil {
		return products.Domain{}, err
	}
	data.ID = d.ID
	return data, nil
}

func (uc Usecase) Delete(id uint) error {
	u, err := uc.data.GetById(id)
	if err != nil {
		return err
	} else if u.ID == 0 {
		return baseErrors.ErrNotFound
	}
	err = uc.data.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
