package Usecase

import (
	"github.com/go-playground/validator/v10"
	"shiva/shiva-auth/internal/products"
	"shiva/shiva-auth/utils/baseErrors"
)

type Usecase struct {
	data     products.Repository
	validate *validator.Validate
}

func NewProductsUsecase(r products.Repository) products.Usecase {
	return &Usecase{
		data:     r,
		validate: validator.New(),
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
	cls, err := uc.data.Create(d)
	if err != nil {
		return products.Domain{}, err
	}
	return cls, nil
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
