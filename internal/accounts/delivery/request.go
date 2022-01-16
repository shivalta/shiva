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

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *Request) ToDomain() accounts.Domain {
	return accounts.Domain{
		Name:       r.Name,
		Email:      r.Email,
		Handphone:  r.Handphone,
		Address:    r.Address,
		Password:   r.Password,
		Repassword: r.Repassword,
	}
}
