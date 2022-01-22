package delivery

import "shiva/shiva-auth/internal/class"

type Response struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	IsPasca bool   `json:"is_pasca"`
	Image   string `json:"image"`
	Slug    string `json:"slug"`
}

func FromDomain(d class.Domain) Response {
	return Response{
		ID:      d.ID,
		Name:    d.Name,
		IsPasca: d.IsPasca,
		Image:   d.ImageUrl,
		Slug:    d.Slug,
	}
}

func FromListDomain(d []class.Domain) (result []Response) {
	for _, v := range d {
		result = append(result, FromDomain(v))
	}
	return result
}
