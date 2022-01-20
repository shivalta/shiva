package delivery

import (
	"shiva/shiva-auth/internal/class"
)

type Request struct {
	Name    string `json:"name,omitempty"`
	IsPasca bool   `json:"is_pasca"`
	Image   string `json:"image,omitempty"`
}

func (r *Request) ToDomain() class.Domain {
	return class.Domain{
		Name:    r.Name,
		IsPasca: r.IsPasca,
		Image:   r.Image,
	}
}
