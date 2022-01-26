package orders

import (
	"mime/multipart"
	"time"
)

type Domain struct {
	ID                uint
	Status            string
	SuccessDateTime   time.Time
	PendingDateTime   time.Time
	FailDateTime      time.Time
	ExpirationPayment time.Time
	UserId            uint
	TotalPrice        int
	TotalTax          float32
	TotalAdmin        int
	AccountNumber     string
	UserValue         string
	UniqueValue       string
	Products          Products
	DetailTransaction DetailTransactionDomain
	BankName          string
	BankCode          string
	Amount            int
	IsLoggedin        bool
}

type DetailTransactionDomain struct {
	ID                        uint
	Sku                       string
	Name                      string
	AdminFee                  int
	Price                     *int
	DetailUniqueValue         string
	DetailUserValue           string
	DetailProductClassName    string
	DetailProductClassImage   string
	DetailProductClassTax     float32
	DetailProductCategoryName string
}

type Products struct {
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
	Checkout(userValue string, productId uint) (Domain, error)
	CreateVA(productId uint, userId uint, bankCode string, userValue string) (Domain, error)
	PaymentChannels() ([]Domain, error)
	GetHistory(userId uint) ([]Domain, error)
	WebhookCreateVA(domain Domain) (Domain, error)
	WebhookPaidVA(domain Domain) (Domain, error)
}

type Repository interface {
	CreateTransaction(domain Domain) (Domain, error)
	UpdateAfterCreateVA(domain Domain) (Domain, error)
	GetHistory(userId uint) ([]Domain, error)
	WebhookCreateVA(domain Domain) (Domain, error)
	WebhookPaidVA(domain Domain) (Domain, error)
}

type XenditRepository interface {
	CreateVA(id string, bankCode string) (Domain, error)
	PaymentChannels() ([]Domain, error)
}

type MockupIoRepository interface {
	GetMockListrik(id string) (Domain, error)
	GetMockPDAM(id string) (Domain, error)
}
