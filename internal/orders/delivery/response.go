package delivery

import "shiva/shiva-auth/internal/orders"

type CheckoutResponse struct {
	UserValue  string `json:"user_value"`
	TotalPrice int    `json:"total_price"`
	IsLoggedin bool   `json:"is_loggedin"`
	Product    ResponseProduct
}

type PaymentMethodResponse struct {
	BankName string `json:"bank_name"`
	BankCode string `json:"bank_code"`
}

type ResponseProduct struct {
	ID                uint               `json:"id"`
	ProductClass      ResponseClass      `json:"product_class,omitempty"`
	ProductCategories ResponseCategories `json:"product_categories,omitempty"`
	Sku               string             `json:"sku"`
	Name              string             `json:"name"`
	AdminFee          int                `json:"admin_fee"`
	Stock             int                `json:"stock,omitempty"`
	Price             int                `json:"price"`
	IsActive          bool               `json:"is_active"`
}

type ResponseClass struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	IsPasca bool   `json:"is_pasca"`
	Image   string `json:"image"`
	Slug    string `json:"slug,omitempty"`
}

type ResponseCategories struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Image string  `json:"image,omitempty"`
	Tax   float32 `json:"tax"`
	Slug  string  `json:"slug,omitempty"`
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
		IsLoggedin: d.IsLoggedin,
		Product: ResponseProduct{
			ID: d.Products.ID,
			ProductClass: ResponseClass{
				ID:      d.Products.ProductClass.ID,
				Name:    d.Products.ProductClass.Name,
				IsPasca: d.Products.ProductClass.IsPasca,
				Image:   d.Products.ProductClass.ImageUrl,
			},
			ProductCategories: ResponseCategories{
				ID:    d.Products.ProductCategory.ID,
				Name:  d.Products.ProductCategory.Name,
				Image: d.Products.ProductCategory.ImageUrl,
				Tax:   d.Products.ProductCategory.Tax,
			},
			Sku:      d.Products.Sku,
			Name:     d.Products.Name,
			AdminFee: d.Products.AdminFee,
			Price:    d.Products.Price,
			IsActive: d.Products.IsActive,
		},
	}
}
