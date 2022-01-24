package orders

import (
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
	DetailTransaction DetailTransactionDomain
	BankName          string
	BankCode          string
	Amount            int
}

type DetailTransactionDomain struct {
	ID                        uint
	Sku                       string
	Name                      string
	AdminFee                  int
	Price                     int
	DetailUniqueValue         string
	DetailProductClassName    string
	DetailProductClassImage   string
	DetailProductClassTax     int
	DetailProductCategoryName string
}

type Usecase interface {
	CheckoutPulsa(userId uint, productId uint) ([]Domain, error)
	CheckoutPDAM(userId uint, productId uint) ([]Domain, error)
	CheckoutListrik(userId uint, productId uint) ([]Domain, error)
	PaymentChannel() (Domain, error)
	CreateVA(productId uint, userId uint, bankCode string) (Domain, error)
	WebhookCreateVA(domain Domain) (Domain, error)
	WebhookPaidVA(domain Domain) (Domain, error)
}

type Repository interface {
	CheckoutPulsa(userId uint, productId uint) ([]Domain, error)
	CheckoutPDAM(userId uint, productId uint) ([]Domain, error)
	CheckoutListrik(userId uint, productId uint) ([]Domain, error)
	PaymentChannel() (Domain, error)
	CreateVA(productId uint, userId uint, bankCode string) (Domain, error)
	WebhookCreateVA(domain Domain) (Domain, error)
	WebhookPaidVA(domain Domain) (Domain, error)
}
