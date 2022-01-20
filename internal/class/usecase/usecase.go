package usecase

import (
	"github.com/go-playground/validator/v10"
	"log"
	"shiva/shiva-auth/helpers/encoder"
	"shiva/shiva-auth/helpers/smtpEmail"
	"shiva/shiva-auth/internal/class"
	"shiva/shiva-auth/utils/baseErrors"
	"shiva/shiva-auth/utils/hash"
)

type accountUsecase struct {
	data     class.Repository
	validate *validator.Validate
}

func NewClassUsecase(r class.Repository) class.Usecase {
	return &accountUsecase{
		data:     r,
		validate: validator.New(),
	}
}

func (uc accountUsecase) Update(user class.Domain) (class.Domain, error) {
	if user.ID != 0 {
		u, err := uc.data.GetById(user.ID)
		if err != nil {
			return class.Domain{}, err
		} else if u.ID == 0 {
			return class.Domain{}, baseErrors.ErrNotFound
		}
	}

	if user.Name == "" {
		return class.Domain{}, baseErrors.ErrUsersNameRequired
	} else if user.Handphone == "" {
		return class.Domain{}, baseErrors.ErrUsersHandphoneRequired
	}
	data, err := uc.data.Update(user)
	if err != nil {
		return class.Domain{}, err
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

func (uc accountUsecase) GetById(id uint) (class.Domain, error) {
	u, err := uc.data.GetById(id)
	if err != nil {
		return class.Domain{}, err
	} else if u.ID == 0 {
		return class.Domain{}, baseErrors.ErrNotFound
	}
	return u, nil
}

func (uc accountUsecase) Create(user class.Domain) (class.Domain, error) {
	if user.Password != user.Repassword {
		return class.Domain{}, baseErrors.ErrInvalidPassword
	}
	u, err := uc.data.GetByEmail(user.Email)
	if err != nil {
		return class.Domain{}, err
	}
	if u.ID != 0 {
		return class.Domain{}, baseErrors.ErrUserEmailUsed
	}
	hashPass, err := hash.HashPassword(user.Password)
	if err != nil {
		return class.Domain{}, nil
	}
	user.Password = hashPass
	res, err := uc.data.Create(user)
	if err != nil {
		return class.Domain{}, err
	}
	url := encoder.EncodeUrlEmailVerify(user.Email)
	bodyEmail := `
		<h2>Hello ` + user.Name + `!</h2>
		Please verify your email with click this link : ` + url +
		`<br><br>Regards,<br>Shiva Admin`
	err = smtpEmail.SendMail([]string{user.Email}, "Email Registration Confirm", bodyEmail)
	if err != nil {
		return class.Domain{}, err
	}
	return res, nil
}

func (uc accountUsecase) GetAll(search string) (data []class.Domain, err error) {
	res, err := uc.data.GetAll(search)
	if err != nil {
		return []class.Domain{}, err
	}
	return res, nil
}
