package model

import (
	"gorm.io/gorm"
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

type ResetPasswords struct {
	gorm.Model
	UserId    int64
	Users     Users     `gorm:"foreignKey:UserId"`
	Status    int32     `gorm:"omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
