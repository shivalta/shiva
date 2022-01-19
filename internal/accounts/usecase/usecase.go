package usecase

import (
	"github.com/go-playground/validator/v10"
	"log"
	"shiva/shiva-auth/cmd/http/middlewares"
	"shiva/shiva-auth/helpers/smtpEmail"
	"shiva/shiva-auth/internal/accounts"
	"shiva/shiva-auth/utils/baseErrors"
	"shiva/shiva-auth/utils/hash"
)

type accountUsecase struct {
	data     accounts.Repository
	validate *validator.Validate
	jwtAuth  *middlewares.ConfigJWT
}

func NewAccountUsecase(r accounts.Repository, jwt *middlewares.ConfigJWT) accounts.Usecase {
	return &accountUsecase{
		data:     r,
		validate: validator.New(),
		jwtAuth:  jwt,
	}
}

func (uc accountUsecase) Login(email string, password string) (string, error) {
	if password == "" {
		return "", baseErrors.ErrUsersPasswordRequired
	} else if email == "" {
		return "", baseErrors.ErrUserEmailRequired
	}
	user, err := uc.data.GetByEmail(email)
	if err != nil {
		return "", err
	}
	if !hash.CheckPassword(password, user.Password) {
		return "", baseErrors.ErrInvalidPassword
	}

	token, err := uc.jwtAuth.GenerateTokenJWT(user.ID, user.IsAdmin)
	if err != nil {
		log.Println(err)
	}
	if token == "" {
		return "", baseErrors.ErrInvalidAuth
	}
	return token, nil
}

func (uc accountUsecase) Update(user accounts.Domain) (accounts.Domain, error) {
	if user.ID != 0 {
		u, err := uc.data.GetById(user.ID)
		if err != nil {
			return accounts.Domain{}, err
		} else if u.ID == 0 {
			return accounts.Domain{}, baseErrors.ErrNotFound
		}
	}

	if user.Name == "" {
		return accounts.Domain{}, baseErrors.ErrUsersNameRequired
	} else if user.Handphone == "" {
		return accounts.Domain{}, baseErrors.ErrUsersHandphoneRequired
	}
	data, err := uc.data.Update(user)
	if err != nil {
		return accounts.Domain{}, err
	}
	return data, nil
}

func (uc accountUsecase) Delete(id uint) error {
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

func (uc accountUsecase) GetById(id uint) (accounts.Domain, error) {
	u, err := uc.data.GetById(id)
	if err != nil {
		return accounts.Domain{}, err
	} else if u.ID == 0 {
		return accounts.Domain{}, baseErrors.ErrNotFound
	}
	return u, nil
}

func (uc accountUsecase) Create(user accounts.Domain) (data accounts.Domain, err error) {
	if user.Password != user.Repassword {
		return accounts.Domain{}, baseErrors.ErrInvalidPassword
	}
	u, err := uc.data.GetByEmail(user.Email)
	if err != nil {
		return accounts.Domain{}, baseErrors.ErrUsersReqNotValid
	}
	if u.ID == 0 {
		return accounts.Domain{}, baseErrors.ErrNotFound
	}
	hashPass, err := hash.HashPassword(user.Password)
	if err != nil {
		return accounts.Domain{}, nil
	}
	user.Password = hashPass
	res, err := uc.data.Create(user)
	if err != nil {
		return accounts.Domain{}, err
	}
	err = smtpEmail.SendMail([]string{user.Email}, "Email Registration Confirm", user.Name+" has register! :)")
	if err != nil {
		return accounts.Domain{}, err
	}
	return res, nil
}

func (uc accountUsecase) GetAll(search string) (data []accounts.Domain, err error) {
	res, err := uc.data.GetAll(search)
	if err != nil {
		return []accounts.Domain{}, err
	}
	return res, nil
}
