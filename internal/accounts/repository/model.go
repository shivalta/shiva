package repository

import (
	"gorm.io/gorm"
	"shiva/shiva-auth/internal/accounts"
	"time"
)

type Admin struct {
	gorm.Model
	Name      string    `gorm:"omitempty"`
	Email     string    `gorm:"omitempty"`
	Password  string    `gorm:"omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Users struct {
	gorm.Model
	Name      string    `gorm:"omitempty"`
	Email     string    `gorm:"omitempty"`
	Handphone string    `gorm:"omitempty"`
	Address   string    `gorm:"omitempty"`
	Password  string    `gorm:"omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

//type ResetPasswords struct {
//	gorm.Model
//	UserId    int64
//	Users     Users     `gorm:"foreignKey:UserId"`
//	Status    int32     `gorm:"omitempty"`
//	CreatedAt time.Time `gorm:"autoCreateTime"`
//	UpdatedAt time.Time `gorm:"autoUpdateTime"`
//}

func (u *Users) UserToDomain() accounts.Domain {
	return accounts.Domain{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Handphone: u.Handphone,
		Address:   u.Address,
		Password:  u.Password,
	}
}

func UserToDomainList(user []Users) (list []accounts.Domain) {
	for _, v := range user {
		list = append(list, v.UserToDomain())
	}
	return list
}
