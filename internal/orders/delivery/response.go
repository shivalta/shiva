package delivery

import (
	"shiva/shiva-auth/internal/orders"
	"time"
)

type TransactionsResponse struct {
	ID                uint                    `json:"id"`
	DetailTransaction DetailTransactionDomain `json:"detail_transaction"`
	Status            string                  `json:"status"`
	SuccessDateTime   time.Time               `json:"success_date_time"`
	PendingDateTime   time.Time               `json:"pending_date_time"`
	FailDateTime      time.Time               `json:"fail_date_time"`
	ExpirationPayment time.Time               `json:"expiration_payment"`
	TotalPrice        int                     `json:"total_price"`
	AccountNumber     string                  `json:"account_number"`
	BankCode          string                  `json:"bank_code"`
}

type DetailTransactionDomain struct {
	ID                        uint    `json:"id"`
	Sku                       string  `json:"sku"`
	Name                      string  `json:"name"`
	AdminFee                  int     `json:"admin_fee"`
	Price                     *int    `json:"price"`
	DetailUniqueValue         string  `json:"detail_unique_value"`
	DetailUserValue           string  `json:"detail_user_value"`
	DetailProductClassName    string  `json:"detail_product_class_name"`
	DetailProductClassImage   string  `json:"detail_product_class_image"`
	DetailProductClassTax     float32 `json:"detail_product_class_tax"`
	DetailProductCategoryName string  `json:"detail_product_category_name"`
}

type CheckoutResponse struct {
	UserValue  string  `json:"user_value"`
	TotalPrice int     `json:"total_price"`
	TotalAdmin int     `json:"total_admin"`
	TotalTax   float32 `json:"total_tax"`
}

type PaymentMethodResponse struct {
	BankName string `json:"bank_name"`
	BankCode string `json:"bank_code"`
}

type CreateVAResponse struct {
	ID                uint      `json:"id"`
	Status            string    `json:"status"`
	TotalPrice        int       `json:"total_price"`
	BankCode          string    `json:"bank_code"`
	AccountName       string    `json:"account_name"`
	AccountNumber     string    `json:"account_number"`
	ExpirationPayment time.Time `json:"expiration_payment"`
}

func FromDomainToTransactionResponse(d orders.Domain) TransactionsResponse {
	return TransactionsResponse{
		ID: d.ID,
		DetailTransaction: DetailTransactionDomain{
			ID:                        d.DetailTransaction.ID,
			Sku:                       d.DetailTransaction.Sku,
			Name:                      d.DetailTransaction.Name,
			AdminFee:                  d.DetailTransaction.AdminFee,
			Price:                     d.DetailTransaction.Price,
			DetailUniqueValue:         d.DetailTransaction.DetailUniqueValue,
			DetailUserValue:           d.DetailTransaction.DetailUserValue,
			DetailProductClassName:    d.DetailTransaction.DetailProductClassName,
			DetailProductClassImage:   d.DetailTransaction.DetailProductClassImage,
			DetailProductClassTax:     d.DetailTransaction.DetailProductClassTax,
			DetailProductCategoryName: d.DetailTransaction.DetailProductCategoryName,
		},
		Status:            d.Status,
		SuccessDateTime:   d.SuccessDateTime,
		PendingDateTime:   d.PendingDateTime,
		FailDateTime:      d.FailDateTime,
		ExpirationPayment: d.ExpirationPayment,
		TotalPrice:        d.TotalPrice,
		AccountNumber:     d.AccountNumber,
		BankCode:          d.BankCode,
	}
}

func FromDomainToTransactionResponseList(d []orders.Domain) []TransactionsResponse {
	t := []TransactionsResponse{}
	for _, v := range d {
		t = append(t, FromDomainToTransactionResponse(v))
	}
	return t
}

func FromDomainToCreateVAResponse(d orders.Domain) CreateVAResponse {
	return CreateVAResponse{
		ID:                d.ID,
		Status:            d.Status,
		TotalPrice:        d.TotalPrice,
		AccountNumber:     d.AccountNumber,
		ExpirationPayment: d.ExpirationPayment,
		BankCode:          d.BankCode,
		AccountName:       d.BankName,
	}
}

func FromDomainPaymentMethod(d orders.Domain) PaymentMethodResponse {
	return PaymentMethodResponse{
		BankName: d.BankName,
		BankCode: d.BankCode,
	}
}

func FromDomainPaymentMethodList(d []orders.Domain) []PaymentMethodResponse {
	var data []PaymentMethodResponse
	for _, v := range d {
		data = append(data, FromDomainPaymentMethod(v))
	}
	return data
}

func FromDomainToCheckout(d orders.Domain) CheckoutResponse {
	return CheckoutResponse{
		UserValue:  d.UserValue,
		TotalPrice: d.TotalPrice,
		TotalAdmin: d.TotalAdmin,
		TotalTax:   d.TotalTax,
	}
}
