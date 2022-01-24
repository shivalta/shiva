package products

import "mime/multipart"

type Domain struct {
	ID                uint
	ProductClassId    uint
	ProductClass      Class
	ProductCategoryId uint
	ProductCategory   Categories
	Sku               string
	Name              string
	AdminFee          int
	Stock             int
	Price             *int
	IsActive          bool
}

type Class struct {
	ID          uint
	Name        string
	IsPasca     bool
	ImageUrl    string
	ImageHeader *multipart.FileHeader
	Slug        string
}

type Categories struct {
	ID             uint
	ProductClassId uint
	ProductClass   Class
	Name           string
	ImageUrl       string
	ImageHeader    *multipart.FileHeader
	Slug           string
	Tax            float32
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
