package dto

import (
	"shiva/shiva-auth/internal/orders"
	"time"
)

type PaymentChannelResponse struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	IsActivated bool   `json:"is_activated"`
}

type CreateVAResponse struct {
	Name           string    `json:"name"`
	AccountNumber  string    `json:"account_number"`
	ExpirationDate time.Time `json:"expiration_date"`
}

func (p *PaymentChannelResponse) PaymentChannelToDomain() orders.Domain {
	return orders.Domain{
		BankName: p.Name,
		BankCode: p.Code,
	}
}

func (p *CreateVAResponse) CreateVAToDomain() orders.Domain {
	return orders.Domain{
		AccountNumber:     p.AccountNumber,
		BankName:          p.Name,
		ExpirationPayment: p.ExpirationDate,
	}
}

func PaymentChannelToDomainList(p []PaymentChannelResponse) []orders.Domain {
	var channel []orders.Domain
	for _, v := range p {
		if v.IsActivated {
			channel = append(channel, v.PaymentChannelToDomain())
		}
	}
	return channel
}
