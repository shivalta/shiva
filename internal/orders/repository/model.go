package repository

import (
	"gorm.io/gorm"
	ruser "shiva/shiva-auth/internal/accounts/repository"
	"shiva/shiva-auth/internal/orders"
	rprod "shiva/shiva-auth/internal/products/repository"
	"time"
)

type Transactions struct {
	gorm.Model
	UserId             uint
	Users              ruser.Users `gorm:"foreignKey:user_id"`
	ProductId          uint
	Products           rprod.Products     `gorm:"foreignKey:product_id"`
	DetailTransactions DetailTransactions `gorm:"foreignKey:transaction_id"`
	Status             string
	SuccessDateTime    time.Time
	PendingDateTime    time.Time
	FailDateTime       time.Time
	ExpirationPayment  time.Time
	TotalPrice         int
	AccountNumber      string
	BankCode           string
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
	DetailProductCategoryTax  float32
	DetailProductCategoryName string
}

func FromDomainToTransaction(d orders.Domain) Transactions {
	return Transactions{
		UserId:            d.UserId,
		ProductId:         d.Products.ID,
		Status:            d.Status,
		SuccessDateTime:   d.SuccessDateTime,
		PendingDateTime:   d.PendingDateTime,
		FailDateTime:      d.FailDateTime,
		ExpirationPayment: d.ExpirationPayment,
		TotalPrice:        d.TotalPrice,
		AccountNumber:     d.AccountNumber,
		BankCode:          d.BankCode,
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
	}
}

func FromDomainToDetail(d orders.Domain) DetailTransactions {
	return DetailTransactions{
		TransactionId:             d.ID,
		Sku:                       d.Products.Sku,
		Name:                      d.Products.Name,
		AdminFee:                  d.Products.AdminFee,
		Price:                     d.Products.Price,
		DetailUniqueUser:          d.UserValue,
		DetailUniqueValue:         d.UniqueValue,
		DetailProductClassName:    d.Products.ProductClass.Name,
		DetailProductClassImage:   d.Products.ProductClass.ImageUrl,
		DetailProductCategoryTax:  d.Products.ProductCategory.Tax,
		DetailProductCategoryName: d.Products.ProductCategory.Name,
	}
}

func (d *Transactions) ToDomain() orders.Domain {
	return orders.Domain{
		ID:                d.ID,
		Status:            d.Status,
		SuccessDateTime:   d.SuccessDateTime,
		PendingDateTime:   d.PendingDateTime,
		FailDateTime:      d.FailDateTime,
		ExpirationPayment: d.ExpirationPayment,
		UserId:            d.UserId,
		TotalPrice:        d.TotalPrice,
		AccountNumber:     d.AccountNumber,
		UserValue:         d.DetailTransactions.DetailUniqueUser,
		UniqueValue:       d.DetailTransactions.DetailUniqueValue,
		Products: orders.Products{
			ID:                d.Products.ID,
			ProductClassId:    d.Products.ID,
			ProductCategoryId: d.Products.ProductCategoryId,
			Sku:               d.Products.Sku,
			Name:              d.Products.Name,
			AdminFee:          d.Products.AdminFee,
			Stock:             d.Products.Stock,
			Price:             *d.Products.Price,
		},
		DetailTransaction: orders.DetailTransactionDomain{
			ID:                        d.DetailTransactions.ID,
			Sku:                       d.DetailTransactions.Sku,
			Name:                      d.DetailTransactions.Name,
			AdminFee:                  d.DetailTransactions.AdminFee,
			Price:                     d.DetailTransactions.Price,
			DetailUniqueValue:         d.DetailTransactions.DetailUniqueValue,
			DetailUserValue:           d.DetailTransactions.DetailUniqueUser,
			DetailProductClassName:    d.DetailTransactions.DetailProductClassName,
			DetailProductClassImage:   d.DetailTransactions.DetailProductClassImage,
			DetailProductClassTax:     d.DetailTransactions.DetailProductCategoryTax,
			DetailProductCategoryName: d.DetailTransactions.DetailProductCategoryName,
		},
		BankCode: d.BankCode,
	}
}
