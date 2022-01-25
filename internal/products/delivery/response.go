package delivery

import (
	"shiva/shiva-auth/internal/products"
)

type Response struct {
	ID                uint               `json:"id"`
	ProductClass      ResponseClass      `json:"product_class,omitempty"`
	ProductCategories ResponseCategories `json:"product_categories,omitempty"`
	Sku               string             `json:"sku"`
	Name              string             `json:"name"`
	AdminFee          int                `json:"admin_fee"`
	Stock             int                `json:"stock"`
	Price             *int               `json:"price"`
	IsActive          bool               `json:"is_active"`
}

type ResponseWithoutForeign struct {
	ID                uint   `json:"id"`
	ProductClassId    uint   `json:"product_class_id"`
	ProductCategoryId uint   `json:"product_category_id"`
	Sku               string `json:"sku"`
	Name              string `json:"name"`
	AdminFee          int    `json:"admin_fee"`
	Stock             int    `json:"stock"`
	Price             *int   `json:"price"`
	IsActive          bool   `json:"is_active"`
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

func FromDomain(d products.Domain) Response {
	return Response{
		ID: d.ID,
		ProductClass: ResponseClass{
			ID:      d.ProductClass.ID,
			Name:    d.ProductClass.Name,
			IsPasca: d.ProductClass.IsPasca,
			Image:   d.ProductClass.ImageUrl,
			Slug:    d.ProductClass.Slug,
		},
		ProductCategories: ResponseCategories{
			ID:    d.ProductCategory.ID,
			Name:  d.ProductCategory.Name,
			Image: d.ProductCategory.ImageUrl,
			Tax:   d.ProductCategory.Tax,
			Slug:  d.ProductCategory.Slug,
		},
		Sku:      d.Sku,
		Name:     d.Name,
		AdminFee: d.AdminFee,
		Stock:    d.Stock,
		Price:    d.Price,
		IsActive: d.IsActive,
	}
}

func FromDomainWithoutForeign(d products.Domain) ResponseWithoutForeign {
	return ResponseWithoutForeign{
		ID:                d.ID,
		ProductClassId:    d.ProductClassId,
		ProductCategoryId: d.ProductCategoryId,
		Sku:               d.Sku,
		Name:              d.Name,
		AdminFee:          d.AdminFee,
		Stock:             d.Stock,
		Price:             d.Price,
		IsActive:          d.IsActive,
	}
}

func FromListDomain(d []products.Domain) (result []Response) {
	for _, v := range d {
		result = append(result, FromDomain(v))
	}
	return result
}
