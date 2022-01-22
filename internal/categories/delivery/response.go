package delivery

import (
	"shiva/shiva-auth/internal/categories"
)

type Response struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Image string  `json:"image,omitempty"`
	Tax   float32 `json:"tax"`
	Slug  string  `json:"slug"`
}

func FromDomain(d categories.Domain) Response {
	return Response{
		ID:    d.ID,
		Name:  d.Name,
		Image: d.ImageUrl,
		Tax:   d.Tax,
		Slug:  d.Slug,
	}
}

func FromListDomain(d []categories.Domain) (result []Response) {
	for _, v := range d {
		result = append(result, FromDomain(v))
	}
	return result
}
