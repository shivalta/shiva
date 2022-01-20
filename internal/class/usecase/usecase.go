package usecase

import (
	"github.com/go-playground/validator/v10"
	"shiva/shiva-auth/internal/class"
	"shiva/shiva-auth/utils/baseErrors"
)

type Usecase struct {
	data     class.Repository
	validate *validator.Validate
}

func NewClassUsecase(r class.Repository) class.Usecase {
	return &Usecase{
		data:     r,
		validate: validator.New(),
	}
}

func (uc Usecase) GetAll(search string, key string) (data []class.Domain, err error) {
	res, err := uc.data.GetAll(search, key)
	if err != nil {
		return []class.Domain{}, err
	}
	return res, nil
}

func (uc Usecase) GetById(id uint) (class.Domain, error) {
	u, err := uc.data.GetById(id)
	if err != nil {
		return class.Domain{}, err
	} else if u.ID == 0 {
		return class.Domain{}, baseErrors.ErrNotFound
	}
	return u, nil
}

func (uc Usecase) Create(d class.Domain) (class.Domain, error) {
	cls, err := uc.data.Create(d)
	if err != nil {
		return class.Domain{}, err
	}
	return cls, nil
}

func (uc Usecase) Update(d class.Domain) (class.Domain, error) {
	data, err := uc.data.Update(d)
	if err != nil {
		return class.Domain{}, err
	}
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
