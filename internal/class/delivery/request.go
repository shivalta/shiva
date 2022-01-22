package delivery

import (
	"mime/multipart"
	"shiva/shiva-auth/internal/class"
)

type Request struct {
	Name    string `json:"name,omitempty" form:"name"`
	IsPasca bool   `json:"is_pasca" form:"is_pasca"`
}

func (r *Request) ToDomain(img *multipart.FileHeader) class.Domain {
	return class.Domain{
		Name:        r.Name,
		IsPasca:     r.IsPasca,
		ImageHeader: img,
	}
}

func (r *Request) ToDomainWithoutImage() class.Domain {
	return class.Domain{
		Name:    r.Name,
		IsPasca: r.IsPasca,
	}
}
