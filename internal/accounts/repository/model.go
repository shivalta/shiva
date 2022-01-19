package repository

import (
	"gorm.io/gorm"
	"shiva/shiva-auth/internal/accounts"
	"time"
)

type Users struct {
	gorm.Model
	Name      string    `gorm:"omitempty"`
	Email     string    `gorm:"omitempty"`
	Handphone string    `gorm:"omitempty"`
	Address   string    `gorm:"omitempty"`
	Password  string    `gorm:"omitempty"`
	IsAdmin   bool      `gorm:"omitempty"`
	IsActive  bool      `gorm:"omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func FromDomain(u accounts.Domain) Users {
	return Users{
		Name:      u.Name,
		Email:     u.Email,
		Handphone: u.Handphone,
		Address:   u.Address,
		Password:  u.Password,
		IsAdmin:   u.IsAdmin,
		IsActive:  u.IsActive,
	}
}

func (u *Users) UserToDomain() accounts.Domain {
	return accounts.Domain{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Handphone: u.Handphone,
		Address:   u.Address,
		Password:  u.Password,
		IsAdmin:   u.IsAdmin,
		IsActive:  u.IsActive,
	}
}

func UserToDomainList(user []Users) (list []accounts.Domain) {
	for _, v := range user {
		list = append(list, v.UserToDomain())
	}
	return list
}
