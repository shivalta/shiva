package repository

import (
	"bytes"
	"encoding/json"
	"net/http"
	"shiva/shiva-auth/internal/orders"
	"shiva/shiva-auth/internal/orders/repository/dto"
)

type XenditAPI struct {
	Client  http.Client
	BaseUrl string
	ApiKey  string
}

func NewXenditAPI(baseUrl string, apiKey string) orders.XenditRepository {
	return &XenditAPI{
		Client:  http.Client{},
		BaseUrl: baseUrl,
		ApiKey:  apiKey,
	}
}

func (api *XenditAPI) PaymentChannels() ([]orders.Domain, error) {
	uri := api.BaseUrl + "/available_virtual_account_banks"
	req, _ := http.NewRequest("GET", uri, nil)
	req.SetBasicAuth(api.ApiKey, "")
	resp, err := api.Client.Do(req)
	if err != nil {
		return []orders.Domain{}, err
	}
	var response []dto.PaymentChannelResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return []orders.Domain{}, err
	}
	return dto.PaymentChannelToDomainList(response), nil
}

func (api *XenditAPI) CreateVA(id string, bankCode string) (orders.Domain, error) {
	uri := api.BaseUrl + "/callback_virtual_accounts"
	body := []byte(`{
		"external_id":"` + id + `",
		"bank_code":"` + bankCode + `",
		"external_id":"PT SHIVA ALTA TBK",
	}`)

	req, _ := http.NewRequest("POST", uri, bytes.NewBuffer(body))
	req.SetBasicAuth(api.ApiKey, "")
	resp, err := api.Client.Do(req)
	if err != nil {
		return orders.Domain{}, err
	}
	var response dto.CreateVAResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return orders.Domain{}, err
	}
	return response.CreateVAToDomain(), nil
}
