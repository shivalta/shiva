package usecase

import (
	"github.com/go-playground/validator/v10"
	"log"
	"shiva/shiva-auth/cmd/http/middlewares"
	"shiva/shiva-auth/helpers/encoder"
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

func (uc accountUsecase) Verify(emailBase64 string, encrypt string) (accounts.Domain, error) {
	email, _ := encoder.DecodeEmailVerify(emailBase64, encrypt)
	if email == "" {
		return accounts.Domain{}, baseErrors.ErrInvalidPayload
	}
	u, err := uc.data.GetByEmail(email)
	if err != nil {
		return accounts.Domain{}, err
	}
	if u.ID == 0 {
		return accounts.Domain{}, baseErrors.ErrUserEmailNotFound
	}
	err = uc.data.UpdateStatus(u.ID, true)
	if err != nil {
		return accounts.Domain{}, err
	}
	bodyEmail := `
		<h2>Hello ` + u.Name + `!</h2>
		Your account has been <font color="green"><b>actived</b></font> :)<br><br>Regards,<br>Shiva Admin`
	err = smtpEmail.SendMail([]string{u.Email}, "SHIVA: Email Verified!", bodyEmail)
	u.IsActive = true
	return u, nil
}

func (uc accountUsecase) Login(email string, password string) (accounts.Domain, string, error) {
	if password == "" {
		return accounts.Domain{}, "", baseErrors.ErrUsersPasswordRequired
	} else if email == "" {
		return accounts.Domain{}, "", baseErrors.ErrUserEmailRequired
	}
	user, err := uc.data.GetByEmail(email)
	if err != nil {
		return accounts.Domain{}, "", err
	}
	if user.IsActive == false {
		return accounts.Domain{}, "", baseErrors.ErrUserNotActive
	}
	if !hash.CheckPassword(password, user.Password) {
		return accounts.Domain{}, "", baseErrors.ErrInvalidPassword
	}

	token, err := uc.jwtAuth.GenerateTokenJWT(user.ID, user.IsAdmin)
	if err != nil {
		log.Println(err)
	}
	if token == "" {
		return accounts.Domain{}, "", baseErrors.ErrInvalidAuth
	}
	return user, token, nil
}

func (uc accountUsecase) Update(user accounts.Domain) (accounts.Domain, error) {
	u, err := uc.data.GetById(user.ID)
	if err != nil {
		return accounts.Domain{}, err
	} else if u.ID == 0 {
		return accounts.Domain{}, baseErrors.ErrNotFound
	}

	if user.Name == "" {
		return accounts.Domain{}, baseErrors.ErrUsersNameRequired
	} else if user.Handphone == "" {
		return accounts.Domain{}, baseErrors.ErrUsersHandphoneRequired
	}

	if user.Password != "" {
		if user.Repassword == user.Password {
			bcryptPassword, _ := hash.HashPassword(user.Password)
			user.Password = bcryptPassword
			data, err := uc.data.UpdateWithPassword(user)
			if err != nil {
				return accounts.Domain{}, err
			}
			return data, nil
		} else {
			return accounts.Domain{}, baseErrors.ErrInvalidPassword
		}
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

func (uc accountUsecase) Create(user accounts.Domain) (accounts.Domain, error) {
	if user.Password != user.Repassword {
		return accounts.Domain{}, baseErrors.ErrInvalidPassword
	}
	u, err := uc.data.GetByEmail(user.Email)
	if err != nil {
		return accounts.Domain{}, err
	}
	if u.ID != 0 {
		return accounts.Domain{}, baseErrors.ErrUserEmailUsed
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
	url := encoder.EncodeUrlEmailVerify(user.Email)
	bodyEmail := `
		<h2>Hello ` + user.Name + `!</h2>
		Please verify your email with click this link : ` + url +
		`<br><br>Regards,<br>Shiva Admin`
	err = smtpEmail.SendMail([]string{user.Email}, "SHIVA: Email Registration Confirm", bodyEmail)
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
