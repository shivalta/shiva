package repository

import (
	"gorm.io/gorm"
	"shiva/shiva-auth/internal/class"
	"strings"
	"time"
)

type ProductClass struct {
	gorm.Model
	Name      string    `gorm:"omitempty"`
	IsPasca   bool      `gorm:"omitempty"`
	Image     string    `gorm:"omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func FromDomain(u class.Domain) ProductClass {
	return ProductClass{
		Name:    u.Name,
		IsPasca: u.IsPasca,
		Image:   u.ImageUrl,
	}
}

func (u *ProductClass) ToDomain() class.Domain {
	return class.Domain{
		ID:       u.ID,
		Name:     u.Name,
		IsPasca:  u.IsPasca,
		ImageUrl: u.Image,
		Slug:     strings.ToLower(strings.ReplaceAll(u.Name, " ", "")),
	}
}

func ToDomainList(user []ProductClass) (list []class.Domain) {
	for _, v := range user {
		list = append(list, v.ToDomain())
	}
	return list
}
