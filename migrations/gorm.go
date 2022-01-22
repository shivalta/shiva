package migrations

import (
	"shiva/shiva-auth/configs/driver"
	accounts "shiva/shiva-auth/internal/accounts/repository"
	categories "shiva/shiva-auth/internal/categories/repository"
	class "shiva/shiva-auth/internal/class/repository"
)

func AutoMigrate() {
	err := driver.Psql.AutoMigrate(
		&accounts.Users{},
		&class.ProductClass{},
		&categories.ProductCategories{},
	)
	if err != nil {
		return
	}
}
