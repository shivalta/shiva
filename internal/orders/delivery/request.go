package delivery

import (
	"time"
)

type RequestCheckout struct {
	UserValue string `json:"user_value"`
	ProductId uint   `json:"product_id"`
}

type RequestPayment struct {
	UserValue string `json:"user_value"`
	ProductId uint   `json:"product_id"`
	BankCode  string `json:"bank_code"`
}

type XenditCallbackRequest struct {
	PaymentId                string    `json:"payment_id"`
	CallbackVirtualAccountId string    `json:"callback_virtual_account_id"`
	OwnerId                  string    `json:"owner_id"`
	ExternalId               string    `json:"external_id"`
	AccountNumber            string    `json:"account_number"`
	BankCode                 string    `json:"bank_code"`
	Amount                   int       `json:"amount"`
	TransactionTimestamp     time.Time `json:"transaction_timestamp"`
	MerchantCode             string    `json:"merchant_code"`
	Id                       string    `json:"id"`
}
