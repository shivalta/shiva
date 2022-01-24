package repository

import "net/http"

type XenditAPI struct {
	Client  http.Client
	BaseUrl string
	ApiKey  string
}

func NewXenditAPI(api XenditAPI) *XenditAPI {
	return &XenditAPI{
		Client:  http.Client{},
		BaseUrl: api.BaseUrl,
		ApiKey:  api.ApiKey,
	}
}

//func (api *XenditAPI) CreateVA(bankCode string, name string) error {
//
//}
