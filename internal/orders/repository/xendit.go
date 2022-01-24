package repository

import (
	"net/http"
	"shiva/shiva-auth/internal/orders"
)

type XenditAPI struct {
	Client  http.Client
	BaseUrl string
	ApiKey  string
}

func NewXenditAPI(client string, baseUrl string, apiKey string) orders.XenditRepository {
	return &XenditAPI{
		Client:  http.Client{},
		BaseUrl: baseUrl,
		ApiKey:  apiKey,
	}
}

func (api *XenditAPI) PaymentChannel() (orders.Domain, error) {
	panic("implement me")
}

func (api *XenditAPI) CreateVA(id string, bankName string, bankCode string) (orders.Domain, error) {
	panic("implement me")
}
