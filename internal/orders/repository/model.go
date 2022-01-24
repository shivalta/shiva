package repository

import (
	"gorm.io/gorm"
	ruser "shiva/shiva-auth/internal/accounts/repository"
	rprod "shiva/shiva-auth/internal/products/repository"
	"time"
)

type Transactions struct {
	gorm.Model
	UserId             uint
	Users              ruser.Users `gorm:"foreignKey:user_id"`
	ProductId          uint
	Products           rprod.Products `gorm:"foreignKey:product_id"`
	DetailTransactions DetailTransactions
	Status             string
	SuccessDateTime    time.Time
	PendingDateTime    time.Time
	FailDateTime       time.Time
	ExpirationPayment  time.Time
	TotalPrice         int
	AccountNumber      string
	CreatedAt          time.Time `gorm:"autoCreateTime"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
}

type DetailTransactions struct {
	gorm.Model
	TransactionId             uint
	Sku                       string
	Name                      string
	AdminFee                  int
	Price                     int
	DetailUniqueUser          string
	DetailUniqueValue         string
	DetailProductClassName    string
	DetailProductClassImage   string
	DetailProductClassTax     int
	DetailProductCategoryName string
}
