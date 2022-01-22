package delivery

import (
	"mime/multipart"
	"shiva/shiva-auth/internal/categories"
)

type Request struct {
	ProductClassId uint    `json:"product_class_id" form:"product_class_id"`
	Name           string  `json:"name,omitempty" form:"name"`
	Tax            float32 `json:"tax" form:"tax"`
}

func (r *Request) ToDomain(img *multipart.FileHeader, classId uint) categories.Domain {
	return categories.Domain{
		ProductClassId: classId,
		ProductClass:   categories.ProductClass{},
		Name:           r.Name,
		ImageHeader:    img,
		Tax:            r.Tax,
	}
}

func (r *Request) ToDomainWithoutImage(classId uint) categories.Domain {
	return categories.Domain{
		ProductClassId: classId,
		ProductClass:   categories.ProductClass{},
		Name:           r.Name,
		Tax:            r.Tax,
	}
}
