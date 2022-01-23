package delivery

import (
	"shiva/shiva-auth/internal/products"
)

type Request struct {
	ProductClassId    uint   `json:"product_class_id"`
	ProductCategoryId uint   `json:"product_category_id"`
	Sku               string `json:"sku"`
	Name              string `json:"name"`
	AdminFee          int    `json:"admin_fee"`
	Stock             int    `json:"stock"`
	Price             int    `json:"price"`
	IsActive          bool   `json:"is_active"`
}

func (r *Request) ToDomain() products.Domain {
	return products.Domain{
		ProductClassId:    r.ProductClassId,
		ProductCategoryId: r.ProductCategoryId,
		Sku:               r.Sku,
		Name:              r.Name,
		AdminFee:          r.AdminFee,
		Stock:             r.Stock,
		Price:             r.Price,
		IsActive:          r.IsActive,
	}
}
