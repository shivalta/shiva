package categories

import (
	"mime/multipart"
)

type Domain struct {
	ID             uint
	ProductClassId uint
	ProductClass   ProductClass
	Name           string
	ImageUrl       string
	ImageHeader    *multipart.FileHeader
	Slug           string
	Tax            float32
}

type ProductClass struct {
	ID          uint
	Name        string
	IsPasca     bool
	ImageUrl    string
	ImageHeader *multipart.FileHeader
	Slug        string
}

type Usecase interface {
	GetAll(search string, key string) ([]Domain, error)
	GetById(id uint) (Domain, error)
	Create(class Domain) (Domain, error)
	Update(class Domain) (Domain, error)
	Delete(id uint) error
}

type Repository interface {
	GetAll(search string, key string) ([]Domain, error)
	GetById(id uint) (Domain, error)
	Create(class Domain) (Domain, error)
	Update(class Domain) (Domain, error)
	UpdateWithoutImage(class Domain) (Domain, error)
	Delete(id uint) error
}
