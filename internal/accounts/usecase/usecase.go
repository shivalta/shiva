package usecase

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"shiva/shiva-auth/internal/accounts"
)

type accountUsecase struct {
	data     accounts.Repository
	validate *validator.Validate
}

func NewAccountUsecase(r accounts.Repository) accounts.Usecase {
	return &accountUsecase{
		data:     r,
		validate: validator.New(),
	}
}

func (uc accountUsecase) Create(user accounts.Domain) (data accounts.Domain, err error) {
	if user.Password != user.Repassword {
		return accounts.Domain{}, errors.New("Password tidak valid")
	}
	res, err := uc.data.Create(user)
	return res, nil
}

func (uc accountUsecase) GetAll(search string) (data []accounts.Domain, err error) {
	res, err := uc.data.GetAll(search)
	if err != nil {
		return []accounts.Domain{}, err
	}
	return res, nil
}
