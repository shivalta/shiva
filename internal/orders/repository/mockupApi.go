package repository

import (
	"net/http"
	"shiva/shiva-auth/internal/orders"
)

type MockupApi struct {
	Client  http.Client
	BaseUrl string
	ApiKey  string
}

func NewMockupApi(baseUrl string, apiKey string) orders.MockupIoRepository {
	return &MockupApi{
		Client:  http.Client{},
		BaseUrl: baseUrl,
		ApiKey:  apiKey,
	}
}

func (m MockupApi) GetName(id string) (orders.Domain, error) {
	panic("implement me")
}
