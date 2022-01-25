package migrations

import (
	"shiva/shiva-auth/configs/driver"
	accounts "shiva/shiva-auth/internal/accounts/repository"
	categories "shiva/shiva-auth/internal/categories/repository"
	class "shiva/shiva-auth/internal/class/repository"
	order "shiva/shiva-auth/internal/orders/repository"
	products "shiva/shiva-auth/internal/products/repository"
)

func AutoMigrate() {
	err := driver.Psql.AutoMigrate(
		&accounts.Users{},
		&class.ProductClass{},
		&categories.ProductCategories{},
		&products.Products{},
		&order.Transactions{},
		&order.DetailTransactions{},
	)
	if err != nil {
		return
	}
}
