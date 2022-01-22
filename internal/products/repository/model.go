package repository

import (
	"gorm.io/gorm"
	rcat "shiva/shiva-auth/internal/categories/repository"
	rclass "shiva/shiva-auth/internal/class/repository"
	"shiva/shiva-auth/internal/products"
	"time"
)

type Products struct {
	gorm.Model
	ProductClassId    uint
	ProductClass      rclass.ProductClass `gorm:"foreignKey:ProductClassId"`
	ProductCategoryId uint
	ProductCategory   rcat.ProductCategories `gorm:"foreignKey:ProductCategoryId"`
	Sku               string
	Name              string
	AdminFee          int
	Stock             int
	Price             int
	IsActive          bool
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
}

func FromDomain(u products.Domain) Products {
	return Products{
		ProductClassId:    u.ProductClassId,
		ProductCategoryId: u.ProductCategoryId,
		Name:              u.Name,
		AdminFee:          u.AdminFee,
		Stock:             u.Stock,
		Price:             u.Price,
		IsActive:          u.IsActive,
	}
}

func (u *Products) ToDomain() products.Domain {
	return products.Domain{
		ID:                u.ID,
		ProductClassId:    u.ProductClassId,
		ProductCategoryId: u.ProductCategoryId,
		Sku:               u.Sku,
		Name:              u.Name,
		AdminFee:          u.AdminFee,
		Stock:             u.Stock,
		Price:             u.Price,
		IsActive:          u.IsActive,
	}
}

func ToDomainList(user []Products) (list []products.Domain) {
	for _, v := range user {
		list = append(list, v.ToDomain())
	}
	return list
}
