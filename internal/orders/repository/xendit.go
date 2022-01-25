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
	body := "{\"external_id\":\"" + id + "\",\"bank_code\":\"" + bankCode + "\",\"name\":\"PT SHIVA ALTA TBK\"}"
	bodyStr := []byte(body)

	req, _ := http.NewRequest("POST", uri, bytes.NewBuffer(bodyStr))
	req.SetBasicAuth(api.ApiKey, "")
	req.Header.Set("Content-Type", "application/json")

	resp, err := api.Client.Do(req)

	if err != nil {
		return orders.Domain{}, err
	}
	response := dto.CreateVAResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return orders.Domain{}, err
	}
	return response.CreateVAToDomain(), nil
}
