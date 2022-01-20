package migrations

import (
	"shiva/shiva-auth/configs/driver"
	accounts "shiva/shiva-auth/internal/accounts/repository"
	class "shiva/shiva-auth/internal/class/repository"
)

func AutoMigrate() {
	err := driver.Psql.AutoMigrate(
		&accounts.Users{},
		&class.ProductClass{},
	)
	if err != nil {
		return
	}
}
