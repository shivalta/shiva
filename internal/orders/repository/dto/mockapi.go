package dto

import (
	"shiva/shiva-auth/internal/orders"
)

type MockapiListrikResponse struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Id       string `json:"id"`
}

type MockapiPDAMResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Bill     int    `json:"bill"`
}

func (m *MockapiListrikResponse) ToDomain() orders.Domain {
	return orders.Domain{
		UserValue: m.Name,
	}
}

func (m *MockapiPDAMResponse) ToDomain() orders.Domain {
	return orders.Domain{
		UserValue:  m.Lastname,
		TotalPrice: m.Bill,
	}
}
