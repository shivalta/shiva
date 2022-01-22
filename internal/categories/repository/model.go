package repository

import (
	"gorm.io/gorm"
	"shiva/shiva-auth/internal/categories"
	rclass "shiva/shiva-auth/internal/class/repository"
	"strings"
	"time"
)

type ProductCategories struct {
	gorm.Model
	ProductClassId uint
	ProductClass   rclass.ProductClass `gorm:"foreignKey:product_class_id"`
	Name           string              `gorm:"omitempty"`
	Image          string              `gorm:"omitempty"`
	Tax            float32             `gorm:"omitempty"`
	CreatedAt      time.Time           `gorm:"autoCreateTime"`
	UpdatedAt      time.Time           `gorm:"autoUpdateTime"`
}

func FromDomain(u categories.Domain) ProductCategories {
	return ProductCategories{
		Name:      u.Name,
		Image:     u.ImageUrl,
		Tax:       u.Tax,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func (u *ProductCategories) ToDomain() categories.Domain {
	return categories.Domain{
		Name:     u.Name,
		ImageUrl: u.Image,
		Slug:     strings.ToLower(strings.ReplaceAll(u.Name, " ", "-")),
		Tax:      u.Tax,
	}
}

func ToDomainList(user []ProductCategories) (list []categories.Domain) {
	for _, v := range user {
		list = append(list, v.ToDomain())
	}
	return list
}
