package delivery

import "shiva/shiva-auth/internal/orders"

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
