package delivery

import "shiva/shiva-auth/internal/accounts"

type Request struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Handphone  string `json:"handphone"`
	Address    string `json:"address"`
	Password   string `json:"password"`
	Repassword string `json:"repassword"`
}

func ToDomain(req Request) accounts.Domain {
	return accounts.Domain{
		Name:       req.Name,
		Email:      req.Email,
		Handphone:  req.Handphone,
		Address:    req.Address,
		Password:   req.Password,
		Repassword: req.Repassword,
	}
}
