package delivery

type RequestCheckout struct {
	UserValue string `json:"user_value"`
	ProductId uint   `json:"product_id"`
}

type RequestPayment struct {
	UserValue string `json:"user_value"`
	ProductId uint   `json:"product_id"`
	BankCode  string `json:"bank_code"`
}
