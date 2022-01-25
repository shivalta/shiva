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
	TotalPrice        int
	AccountNumber     string
	UserValue         string
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
	Price                     int
	DetailUniqueValue         string
	DetailUserValue           string
	DetailProductClassName    string
	DetailProductClassImage   string
	DetailProductClassTax     int
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
	Price             int
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
	CheckoutPulsa(userValue string, productId uint, isLoggedIn bool) (Domain, error)
	CheckoutPDAM(userValue string, productId uint, isLoggedIn bool) (Domain, error)
	CheckoutListrik(userValue string, productId uint, isLoggedIn bool) (Domain, error)
	CreateVA(productId uint, userId uint, bankCode string) (Domain, error)
	PaymentChannels() ([]Domain, error)
	WebhookCreateVA(domain Domain) (Domain, error)
	WebhookPaidVA(domain Domain) (Domain, error)
}

type Repository interface {
	CheckoutPulsa(userId uint, productId uint) (Domain, error)
	CheckoutPDAM(userId uint, productId uint) (Domain, error)
	CheckoutListrik(userId uint, productId uint) (Domain, error)
	CreateTransaction(productId uint, userId uint, bankCode string) (Domain, error)
	WebhookCreateVA(domain Domain) (Domain, error)
	WebhookPaidVA(domain Domain) (Domain, error)
}

type XenditRepository interface {
	CreateVA(id string, bankName string, bankCode string) (Domain, error)
	PaymentChannels() ([]Domain, error)
}

type MockupIoRepository interface {
	GetName(id string) (Domain, error)
}
