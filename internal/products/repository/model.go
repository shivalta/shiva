package repository

import (
	"gorm.io/gorm"
	rcat "shiva/shiva-auth/internal/categories/repository"
	rclass "shiva/shiva-auth/internal/class/repository"
	"shiva/shiva-auth/internal/products"
	"strings"
	"time"
)

type Products struct {
	gorm.Model
	ProductClassId    uint
	ProductClass      rclass.ProductClass `gorm:"foreignKey:product_class_id"`
	ProductCategoryId uint
	ProductCategory   rcat.ProductCategories `gorm:"foreignKey:product_category_id"`
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
		Sku:               u.Sku,
		Name:              u.Name,
		AdminFee:          u.AdminFee,
		Stock:             u.Stock,
		Price:             u.Price,
		IsActive:          u.IsActive,
	}
}

func (u *Products) ToDomain() products.Domain {
	return products.Domain{
		ID: u.ID,
		ProductClass: products.Class{
			ID:       u.ProductClass.ID,
			Name:     u.ProductClass.Name,
			IsPasca:  u.ProductClass.IsPasca,
			ImageUrl: u.ProductClass.Image,
			Slug:     strings.ToLower(strings.ReplaceAll(u.ProductClass.Name, " ", "-")),
		},
		ProductCategory: products.Categories{
			ID:             u.ProductCategory.ID,
			ProductClassId: u.ProductCategory.ProductClassId,
			Name:           u.ProductCategory.Name,
			ImageUrl:       u.ProductCategory.Image,
			Tax:            u.ProductCategory.Tax,
			Slug:           strings.ToLower(strings.ReplaceAll(u.ProductCategory.Name, " ", "-")),
		},
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
