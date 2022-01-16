package delivery

import "shiva/shiva-auth/internal/accounts"

type Response struct {
	ID        uint   `json:"id,omitempty"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Handphone string `json:"handphone"`
	Address   string `json:"address"`
}

func FromDomain(d accounts.Domain) Response {
	return Response{
		ID:        d.ID,
		Name:      d.Name,
		Email:     d.Email,
		Handphone: d.Handphone,
		Address:   d.Address,
	}
}

func FromListDomain(d []accounts.Domain) (result []Response) {
	for _, v := range d {
		result = append(result, FromDomain(v))
	}
	return result
}
