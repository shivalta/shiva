package dto

import (
	"shiva/shiva-auth/internal/orders"
)

type PaymentChannelResponse struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	IsActivated bool   `json:"is_activated"`
}

func (p *PaymentChannelResponse) PaymentChannelToDomain() orders.Domain {
	return orders.Domain{
		BankName: p.Name,
		BankCode: p.Code,
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
