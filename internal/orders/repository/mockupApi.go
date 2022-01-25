package repository

import (
	"encoding/json"
	"net/http"
	"shiva/shiva-auth/internal/orders"
	"shiva/shiva-auth/internal/orders/repository/dto"
)

type MockupApi struct {
	Client  http.Client
	BaseUrl string
}

func NewMockupApi(baseUrl string) orders.MockupIoRepository {
	return &MockupApi{
		Client:  http.Client{},
		BaseUrl: baseUrl,
	}
}

func (api MockupApi) GetMockListrik(id string) (orders.Domain, error) {
	uri := api.BaseUrl + "/api/v1/listrik/" + id
	req, _ := http.NewRequest("GET", uri, nil)
	resp, err := api.Client.Do(req)
	if err != nil {
		return orders.Domain{}, err
	}
	var response dto.MockapiListrikResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return orders.Domain{}, err
	}
	return response.ToDomain(), nil
}

func (api MockupApi) GetMockPDAM(id string) (orders.Domain, error) {
	uri := api.BaseUrl + "/api/v1/pdam/" + id
	req, _ := http.NewRequest("GET", uri, nil)
	resp, err := api.Client.Do(req)
	if err != nil {
		return orders.Domain{}, err
	}
	var response dto.MockapiPDAMResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return orders.Domain{}, err
	}
	return response.ToDomain(), nil
}
